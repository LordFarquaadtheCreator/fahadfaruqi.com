#!/usr/bin/env python3
"""Upload static/interp.wasm and its matching wasm_exec.js to the CDN.

The interpreter is not shipped in the build: the site fetches both files from the
Cloudflare CDN (R2 bucket `assets`, www/ prefix). wasm_exec.js is taken from the same
Go toolchain that built the wasm, so the pair can never drift apart.

Credentials come from the environment:

    R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... ./scripts/publish-wasm.py
"""

import os
import subprocess
import sys

import boto3

BUCKET = "assets"
PREFIX = "www"
ENDPOINT = "https://f4c3ab68be5b33d51bf06ea35074f2c6.r2.cloudflarestorage.com"
# Cloudflare raises the browser TTL of cacheable types below its zone-level Browser
# Cache TTL to 14400s, so wasm_exec.js asks for that outright instead of a value the
# CDN silently overrides. .wasm is not in Cloudflare's cached-extension list, so the
# interpreter is fetched per visit and 300s is honoured as written.
WASM_CACHE_CONTROL = "public, max-age=300"
EXEC_JS_CACHE_CONTROL = "public, max-age=14400"
ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))


def goroot():
    return subprocess.run(
        ["go", "env", "GOROOT"], capture_output=True, text=True, check=True
    ).stdout.strip()


def main():
    access_key = os.environ.get("R2_ACCESS_KEY_ID")
    secret_key = os.environ.get("R2_SECRET_ACCESS_KEY")
    if not access_key or not secret_key:
        sys.exit("R2_ACCESS_KEY_ID and R2_SECRET_ACCESS_KEY must be set")

    wasm = os.path.join(ROOT, "interp.wasm")
    if not os.path.isfile(wasm):
        sys.exit(f"not built: {wasm} — run ./build-wasm.sh first")

    exec_js = os.path.join(goroot(), "lib", "wasm", "wasm_exec.js")
    if not os.path.isfile(exec_js):
        sys.exit(f"wasm_exec.js not found at {exec_js}")

    s3 = boto3.client(
        "s3",
        endpoint_url=ENDPOINT,
        region_name="auto",
        aws_access_key_id=access_key,
        aws_secret_access_key=secret_key,
    )

    for local, key, content_type, cache_control in (
        (wasm, f"{PREFIX}/interp.wasm", "application/wasm", WASM_CACHE_CONTROL),
        (exec_js, f"{PREFIX}/wasm_exec.js", "text/javascript", EXEC_JS_CACHE_CONTROL),
    ):
        s3.upload_file(
            local,
            BUCKET,
            key,
            ExtraArgs={
                "ContentType": content_type,
                "CacheControl": cache_control,
                "Metadata": {"site": "fahadfaruqi.com", "source": os.path.basename(local)},
            },
        )
        head = s3.head_object(Bucket=BUCKET, Key=key)
        if head["ContentLength"] != os.path.getsize(local):
            sys.exit(f"size mismatch for {key}: {head['ContentLength']} stored")
        print(
            f"{key}: {head['ContentLength']} bytes, {head['ContentType']}, "
            f"{head.get('CacheControl')} ({head.get('Metadata')})"
        )


if __name__ == "__main__":
    main()
