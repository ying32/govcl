set GOARCH=amd64
set CGO_ENABLED=1
go build -i -ldflags="-s -w" -tags memorydll
pause