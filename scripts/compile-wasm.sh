#!/usr/bin/env bash

cd wasm
CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -o web/coca.wasm -v