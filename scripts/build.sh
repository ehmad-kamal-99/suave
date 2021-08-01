#!/bin/bash

mkdir -p "bin"

go build -o "bin/suave" cmd/suave-server/main.go
