#!/bin/bash
tinygo build -o hello-world.wasm -scheduler=none --no-debug -target wasi ./main.go


