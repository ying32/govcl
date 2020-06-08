set GOARCH=386
go build -i -ldflags="-s -w" -tags memorydll
pause