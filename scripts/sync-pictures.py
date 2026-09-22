#!/usr/bin/env python3
"""Upload pictures to the CDN and sync the ~/Pictures children in data.json.

Originals and thumbnails are not in the repo: they live in the R2 bucket behind the
CDN, so data.json only carries the relative paths and the Go wasm prefixes them with
the asset base at runtime. Credentials come from the environment:

    R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... ./scripts/sync-pictures.py [source-dir]

The source directory defaults to ~/Pictures and is never modified.
"""

import json
import os
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
IMAGE_EXTS = {".png", ".jpg", ".jpeg", ".gif", ".webp"}
MIME_TYPES = {
    ".png": "image/png",
    ".jpg": "image/jpeg",
    ".jpeg": "image/jpeg",
    ".gif": "image/gif",
    ".webp": "image/webp",
}


def content_type(name):
    return MIME_TYPES.get(os.path.splitext(name)[1].lower(), "application/octet-stream")


def dimensions(path):
    out = subprocess.run(
        ["sips", "-g", "pixelWidth", "-g", "pixelHeight", path],
        capture_output=True, text=True, check=False,
    )
    w = h = 0
    for line in out.stdout.splitlines():
        if line.startswith("pixelWidth"):
            w = int(line.split(":")[1].strip())
        elif line.startswith("pixelHeight"):
            h = int(line.split(":")[1].strip())
    return w, h


def main():
    access_key = os.environ.get("R2_ACCESS_KEY_ID")
    secret_key = os.environ.get("R2_SECRET_ACCESS_KEY")
    if not access_key or not secret_key:
        sys.exit("R2_ACCESS_KEY_ID and R2_SECRET_ACCESS_KEY must be set")

    source = os.path.abspath(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_SOURCE
    if not os.path.isdir(source):
        sys.exit(f"pictures dir not found: {source}")

    image_files = sorted(
        f for f in os.listdir(source)
        if os.path.splitext(f)[1].lower() in IMAGE_EXTS
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
    with tempfile.TemporaryDirectory() as tmp:
        for name in image_files:
            src = os.path.join(source, name)
            thumb = os.path.join(tmp, name)
            subprocess.run(
                ["sips", "-z", str(THUMB_SIZE), str(THUMB_SIZE), src, "--out", thumb],
                capture_output=True, check=False,
            )
            w, h = dimensions(src)

            for local, key, label in (
                (src, f"{KEY_PREFIX}/{name}", "original"),
                (thumb, f"{KEY_PREFIX}/thumbs/{name}", "thumb"),
            ):
                s3.upload_file(local, BUCKET, key, ExtraArgs={
                    "ContentType": content_type(name),
                    "CacheControl": CACHE_CONTROL,
                    "Metadata": {"site": "fahadfaruqi.com",
                                 "source": f"static/images/pictures/{'thumbs/' if label == 'thumb' else ''}{name}"},
                })
                head = s3.head_object(Bucket=BUCKET, Key=key)
                if head["ContentLength"] != os.path.getsize(local):
                    sys.exit(f"size mismatch for {key}: {head['ContentLength']} stored")

            data[f"~/Pictures/{name}"] = {
                "type": "file",
                "content": "",
                "imagePath": f"/images/pictures/{name}",
                "thumbPath": f"/images/pictures/thumbs/{name}",
                "imgWidth": w,
                "imgHeight": h,
                "mimeType": MIME_TYPES.get(os.path.splitext(name)[1].lower(), ""),
                "hidden": False,
            }
            children.append(f"~/Pictures/{name}")

    if "~/Pictures" not in data:
        sys.exit("~/Pictures not found in data.json")
    data["~/Pictures"]["children"] = children

    with open(DATA_JSON, "w") as fh:
        json.dump(data, fh, indent=2)
        fh.write("\n")

    print(f"uploaded and synced {len(children)} image(s) from {source}")


if __name__ == "__main__":
    main()
