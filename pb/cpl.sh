#!/usr/bin/env bash
export PATH=$PATH:/home/jack/go/bin
rm -f ./pbcpl/*
protoc -I=./pbsrc --go_out=./pbcpl ./pbsrc/*.proto
