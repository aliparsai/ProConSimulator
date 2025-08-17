@echo off
echo "Building for Windows..."
go build -o app.exe
echo "Build complete. Run with app.exe -type=server or app.exe -type=client"
