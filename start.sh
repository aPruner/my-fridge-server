#!/bin/bash

if [ ! -d "build" ]
then
  mkdir build
fi
go build -v -o build/ ./... && ./build/my-fridge-server