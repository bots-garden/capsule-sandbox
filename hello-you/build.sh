#!/bin/bash
tinygo build -o hello-you.wasm -scheduler=none --no-debug -target wasi ./main.go


