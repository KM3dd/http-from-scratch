#!/bin/sh


set -e # Exit on failure

go build -o /tmp/codecrafters-build-http-server-go app/*.go
