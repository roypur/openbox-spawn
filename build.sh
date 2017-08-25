#!/bin/bash
mkdir -p bin
go build -o bin/wrapper wrapper.go
go build -o bin/spawn spawn.go
