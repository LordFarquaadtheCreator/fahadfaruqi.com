#!/usr/bin/env python3
"""Scan static/images/pictures/ and sync ~/Pictures children in data.json."""

import json
import os
import sys

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
PICTURES_DIR = os.path.join(ROOT, "static", "images", "pictures")
DATA_JSON = os.path.join(ROOT, "interpreter", "world", "data.json")

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

    # Scan for image files
    image_files = sorted(
        f for f in os.listdir(PICTURES_DIR)
        if os.path.splitext(f)[1].lower() in IMAGE_EXTS
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
        data[node_path] = {
            "type": "file",
            "content": "",
            "imagePath": f"/images/pictures/{name}",
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

    print(f"synced {len(children)} image(s) to ~/Pictures")


if __name__ == "__main__":
    main()
