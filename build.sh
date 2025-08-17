#!/bin/bash
echo "Building for Linux..."
go build -o app
echo "Build complete. Run with ./app -type=server or ./app -type=client"
