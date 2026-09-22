#!/bin/bash

cd interpreter
GOOS=js GOARCH=wasm go build \
  -ldflags "-X interpreter/assets.Base=${ASSETS_BASE:-https://assets.fahadfaruqi.com/www}" \
  -o ../interp.wasm
echo "built interp.wasm — publish it with scripts/publish-wasm.py"
