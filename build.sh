#!/bin/sh
set -xe

GOOS=linux GOARCH=arm go build -o tuodan .
