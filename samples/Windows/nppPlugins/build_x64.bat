@echo off
echo set env
set GOARCH=amd64
set CGO_ENABLED=1
set PATH=%PATH%;%mingw64%;%mingw64%\bin

echo build 
go build -i -ldflags="-s -w" -buildmode=c-shared  -o nppPlugins.dll
 
pause
