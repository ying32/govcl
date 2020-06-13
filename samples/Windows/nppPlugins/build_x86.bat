@echo off
echo set env
set GOARCH=386
set CGO_ENABLED=1
set mingw64=%MinGW32%
set PATH=%PATH%;%MinGW32%;%MinGW32%\bin

echo build 
go build -i -ldflags="-s -w" -buildmode=c-shared  -o nppPlugins.dll

pause
