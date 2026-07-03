#!/bin/bash

cd interpreter
GOOS=js GOARCH=wasm go build -o ../static/interp.wasm
