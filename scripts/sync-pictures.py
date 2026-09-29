#!/usr/bin/env python3
"""Upload pictures to the CDN and sync the ~/Pictures children in data.json.

Pictures are not in the repo: they live in the R2 bucket behind the CDN, so data.json
only carries the relative paths and the Go wasm prefixes them with the asset base at
runtime. Each image becomes three objects:

    <prefix>/<name>            the master, as it came off the camera
    <prefix>/thumbs/<name>     100px thumbnail for Finder icons (sips)
    <prefix>/compressed/<stem>.avif   the copy the simulation renders (ffmpeg, AV1)

Credentials come from the environment:

    R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... ./scripts/sync-pictures.py [source-dir]

The source directory defaults to ~/Pictures and is never modified. ffmpeg and ffprobe
must be on PATH; sips does the thumbnails.
"""

import json
import os
import shutil
import subprocess
import sys
import tempfile

import boto3

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
DATA_JSON = os.path.join(ROOT, "interpreter", "world", "data.json")
DEFAULT_SOURCE = os.path.expanduser("~/Pictures")

BUCKET = "assets"
ENDPOINT = "https://f4c3ab68be5b33d51bf06ea35074f2c6.r2.cloudflarestorage.com"
KEY_PREFIX = "www/images/pictures"
CACHE_CONTROL = "public, max-age=604800"

THUMB_SIZE = 100
COMPRESSED_MAX = 1600  # longest side of the compressed copy
COMPRESSED_CRF = 22
COMPRESSED_PRESET = 6
IMAGE_EXTS = {".png", ".jpg", ".jpeg", ".gif", ".webp"}
# An animated gif has no single-frame compressed copy: it keeps pointing at its master.
STILL_EXTS = IMAGE_EXTS - {".gif"}
MIME_TYPES = {
    ".png": "image/png",
    ".jpg": "image/jpeg",
    ".jpeg": "image/jpeg",
    ".gif": "image/gif",
    ".webp": "image/webp",
    ".avif": "image/avif",
}


def ext(name):
    return os.path.splitext(name)[1].lower()


def content_type(name):
    return MIME_TYPES.get(ext(name), "application/octet-stream")


def dimensions(path):
    """Natural pixel size with the EXIF rotation applied — the size a browser shows."""
    out = subprocess.run(
        ["ffprobe", "-v", "error", "-select_streams", "v:0",
         "-show_entries", "stream=width,height", "-show_entries", "side_data=rotation",
         "-of", "default=nw=1", path],
        capture_output=True, text=True, check=False,
    )
    w = h = rotation = 0
    for line in out.stdout.splitlines():
        key, _, value = line.partition("=")
        try:
            if key == "width":
                w = int(value)
            elif key == "height":
                h = int(value)
            elif key == "rotation":
                rotation = int(float(value))
        except ValueError:
            pass
    if rotation % 180 != 0:  # a quarter turn: the displayed frame is the other way round
        w, h = h, w
    return w, h


def thumbnail(source, name, directory):
    thumb = os.path.join(directory, name)
    subprocess.run(
        ["sips", "-z", str(THUMB_SIZE), str(THUMB_SIZE), source, "--out", thumb],
        capture_output=True, check=False,
    )
    return thumb


def compress(source, directory, stem):
    """Write the AVIF copy the simulation renders. Returns the path, or None."""
    out = os.path.join(directory, stem + ".avif")
    fit = f"scale='min({COMPRESSED_MAX},iw)':'min({COMPRESSED_MAX},ih)':force_original_aspect_ratio=decrease"
    result = subprocess.run(
        ["ffmpeg", "-v", "error", "-y", "-i", source, "-vf", fit,
         "-c:v", "libsvtav1", "-crf", str(COMPRESSED_CRF),
         "-preset", str(COMPRESSED_PRESET), "-frames:v", "1", out],
        capture_output=True, text=True, check=False,
    )
    if result.returncode != 0 or not os.path.isfile(out):
        sys.exit(f"ffmpeg failed for {source}: {result.stderr.strip()}")
    return out


def put(s3, local, key):
    s3.upload_file(local, BUCKET, key, ExtraArgs={
        "ContentType": content_type(key),
        "CacheControl": CACHE_CONTROL,
        "Metadata": {"site": "fahadfaruqi.com", "source": os.path.basename(local)},
    })
    head = s3.head_object(Bucket=BUCKET, Key=key)
    if head["ContentLength"] != os.path.getsize(local):
        sys.exit(f"size mismatch for {key}: {head['ContentLength']} stored")
    return head["ContentLength"]


def main():
    access_key = os.environ.get("R2_ACCESS_KEY_ID")
    secret_key = os.environ.get("R2_SECRET_ACCESS_KEY")
    if not access_key or not secret_key:
        sys.exit("R2_ACCESS_KEY_ID and R2_SECRET_ACCESS_KEY must be set")

    for tool in ("sips", "ffmpeg", "ffprobe"):
        if not shutil.which(tool):
            sys.exit(f"{tool} not found on PATH")

    source = os.path.abspath(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_SOURCE
    if not os.path.isdir(source):
        sys.exit(f"pictures dir not found: {source}")

    image_files = sorted(
        f for f in os.listdir(source)
        if ext(f) in IMAGE_EXTS
        and os.path.isfile(os.path.join(source, f))
    )

    s3 = boto3.client(
        "s3", endpoint_url=ENDPOINT, region_name="auto",
        aws_access_key_id=access_key, aws_secret_access_key=secret_key,
    )

    with open(DATA_JSON) as fh:
        data = json.load(fh)

    for key in [k for k in data if k.startswith("~/Pictures/")]:
        del data[key]

    children = []
    master_bytes = compressed_bytes = 0
    with tempfile.TemporaryDirectory() as tmp:
        for name in image_files:
            src = os.path.join(source, name)
            stem = os.path.splitext(name)[0]
            w, h = dimensions(src)

            master_bytes += put(s3, src, f"{KEY_PREFIX}/{name}")
            put(s3, thumbnail(src, name, tmp), f"{KEY_PREFIX}/thumbs/{name}")

            compressed_path = ""
            if ext(name) in STILL_EXTS:
                copy = compress(src, tmp, stem)
                compressed_bytes += put(s3, copy, f"{KEY_PREFIX}/compressed/{stem}.avif")
                compressed_path = f"/images/pictures/compressed/{stem}.avif"

            data[f"~/Pictures/{name}"] = {
                "type": "file",
                "content": "",
                "imagePath": f"/images/pictures/{name}",
                "thumbPath": f"/images/pictures/thumbs/{name}",
                "compressedPath": compressed_path,
                "imgWidth": w,
                "imgHeight": h,
                "mimeType": MIME_TYPES.get(ext(name), ""),
                "hidden": False,
            }
            children.append(f"~/Pictures/{name}")

    if "~/Pictures" not in data:
        sys.exit("~/Pictures not found in data.json")
    data["~/Pictures"]["children"] = children

    with open(DATA_JSON, "w") as fh:
        json.dump(data, fh, indent=2)
        fh.write("\n")

    print(
        f"uploaded and synced {len(children)} image(s) from {source}: "
        f"{master_bytes / 1048576:.1f} MB of masters, "
        f"{compressed_bytes / 1048576:.1f} MB of compressed copies"
    )


if __name__ == "__main__":
    main()
