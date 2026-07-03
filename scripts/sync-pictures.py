#!/usr/bin/env python3
"""Scan static/images/pictures/ and sync ~/Pictures children in data.json."""

import json
import os
import subprocess
import sys

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
PICTURES_DIR = os.path.join(ROOT, "static", "images", "pictures")
THUMBS_DIR = os.path.join(PICTURES_DIR, "thumbs")
DATA_JSON = os.path.join(ROOT, "interpreter", "world", "data.json")

THUMB_SIZE = 100

IMAGE_EXTS = {".png", ".jpg", ".jpeg", ".gif", ".webp"}

MIME_TYPES = {
    ".png": "image/png",
    ".jpg": "image/jpeg",
    ".jpeg": "image/jpeg",
    ".gif": "image/gif",
    ".webp": "image/webp",
}


def main():
    if not os.path.isdir(PICTURES_DIR):
        print(f"pictures dir not found: {PICTURES_DIR}", file=sys.stderr)
        sys.exit(1)

    os.makedirs(THUMBS_DIR, exist_ok=True)

    # Scan for image files (exclude thumbs dir)
    image_files = sorted(
        f for f in os.listdir(PICTURES_DIR)
        if os.path.splitext(f)[1].lower() in IMAGE_EXTS
        and os.path.isfile(os.path.join(PICTURES_DIR, f))
    )

    with open(DATA_JSON, "r") as fh:
        data = json.load(fh)

    # Remove stale ~/Pictures/* nodes
    stale = [k for k in data if k.startswith("~/Pictures/")]
    for k in stale:
        del data[k]

    # Add new image nodes and build children list
    children = []
    for name in image_files:
        node_path = f"~/Pictures/{name}"
        ext = os.path.splitext(name)[1].lower()
        src = os.path.join(PICTURES_DIR, name)
        thumb = os.path.join(THUMBS_DIR, name)

        # Generate thumbnail
        subprocess.run(
            ["sips", "-z", str(THUMB_SIZE), str(THUMB_SIZE), src, "--out", thumb],
            capture_output=True, check=False,
        )

        # Get dimensions
        w, h = 0, 0
        out = subprocess.run(
            ["sips", "-g", "pixelWidth", "-g", "pixelHeight", src],
            capture_output=True, text=True, check=False,
        )
        for line in out.stdout.splitlines():
            if line.startswith("pixelWidth"):
                w = int(line.split(":")[1].strip())
            elif line.startswith("pixelHeight"):
                h = int(line.split(":")[1].strip())

        data[node_path] = {
            "type": "file",
            "content": "",
            "imagePath": f"/images/pictures/{name}",
            "thumbPath": f"/images/pictures/thumbs/{name}",
            "imgWidth": w,
            "imgHeight": h,
            "mimeType": MIME_TYPES.get(ext, ""),
            "hidden": False,
        }
        children.append(node_path)

    # Update ~/Pictures children
    if "~/Pictures" not in data:
        print("~/Pictures not found in data.json", file=sys.stderr)
        sys.exit(1)
    data["~/Pictures"]["children"] = children

    with open(DATA_JSON, "w") as fh:
        json.dump(data, fh, indent=2)
        fh.write("\n")

    print(f"synced {len(children)} image(s) to ~/Pictures ({len(children)} thumbnails generated)")


if __name__ == "__main__":
    main()
