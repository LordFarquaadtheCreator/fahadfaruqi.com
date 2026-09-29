#!/usr/bin/env python3
"""Write the compressed copy of each desktop wallpaper to the CDN.

The wallpapers are not in the repo: the masters sit in the R2 bucket behind the CDN at
`www/wallpapers/wallpaper.jpg` and `www/wallpapers/wallpaper-dark.jpg`. The desktop
background is fetched on every visit, so `interpreter/desktop/desktop.go` points at the
compressed copy instead:

    www/wallpapers/compressed/wallpaper.avif
    www/wallpapers/compressed/wallpaper-dark.avif

The masters are read back from the bucket, encoded, and left in place.

    R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... ./scripts/sync-wallpapers.py

Replacing a wallpaper means replacing the master in the bucket (`www/wallpapers/…jpg`),
then running this script. ffmpeg must be on PATH.
"""

import os
import shutil
import subprocess
import sys
import tempfile

import boto3

BUCKET = "assets"
ENDPOINT = "https://f4c3ab68be5b33d51bf06ea35074f2c6.r2.cloudflarestorage.com"
KEY_PREFIX = "www/wallpapers"
CACHE_CONTROL = "public, max-age=604800"

NAMES = ("wallpaper.jpg", "wallpaper-dark.jpg")
COMPRESSED_MAX = 3200  # longest side; covers a retina laptop screen without upscaling
COMPRESSED_CRF = 22
COMPRESSED_PRESET = 6


def compress(source, out):
    fit = f"scale='min({COMPRESSED_MAX},iw)':'min({COMPRESSED_MAX},ih)':force_original_aspect_ratio=decrease"
    result = subprocess.run(
        ["ffmpeg", "-v", "error", "-y", "-i", source, "-vf", fit,
         "-c:v", "libsvtav1", "-crf", str(COMPRESSED_CRF),
         "-preset", str(COMPRESSED_PRESET), "-frames:v", "1", out],
        capture_output=True, text=True, check=False,
    )
    if result.returncode != 0 or not os.path.isfile(out):
        sys.exit(f"ffmpeg failed for {source}: {result.stderr.strip()}")


def main():
    access_key = os.environ.get("R2_ACCESS_KEY_ID")
    secret_key = os.environ.get("R2_SECRET_ACCESS_KEY")
    if not access_key or not secret_key:
        sys.exit("R2_ACCESS_KEY_ID and R2_SECRET_ACCESS_KEY must be set")

    if not shutil.which("ffmpeg"):
        sys.exit("ffmpeg not found on PATH")

    s3 = boto3.client(
        "s3", endpoint_url=ENDPOINT, region_name="auto",
        aws_access_key_id=access_key, aws_secret_access_key=secret_key,
    )

    with tempfile.TemporaryDirectory() as tmp:
        for name in NAMES:
            stem = os.path.splitext(name)[0]
            master_key = f"{KEY_PREFIX}/{name}"
            master = os.path.join(tmp, name)
            s3.download_file(BUCKET, master_key, master)

            copy = os.path.join(tmp, stem + ".avif")
            compress(master, copy)

            key = f"{KEY_PREFIX}/compressed/{stem}.avif"
            s3.upload_file(copy, BUCKET, key, ExtraArgs={
                "ContentType": "image/avif",
                "CacheControl": CACHE_CONTROL,
                "Metadata": {"site": "fahadfaruqi.com", "source": name},
            })
            head = s3.head_object(Bucket=BUCKET, Key=key)
            if head["ContentLength"] != os.path.getsize(copy):
                sys.exit(f"size mismatch for {key}: {head['ContentLength']} stored")

            print(
                f"{key}: {head['ContentLength'] / 1024:.0f} KB "
                f"(master {os.path.getsize(master) / 1048576:.1f} MB, "
                f"{head['ContentType']}, {head.get('CacheControl')})"
            )


if __name__ == "__main__":
    main()
