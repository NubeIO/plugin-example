#!/bin/bash

# Step 1: Build and move files
go build -o math-plugin *.go && mv -f math-plugin /home/aidan/code/go/rubix-rx/data/extensions/

# Step 2: Change directory
(cd /home/aidan/code/go/rubix-rx &&  go run main.go --id=R-1 --port=1770 --grpc=9090)


