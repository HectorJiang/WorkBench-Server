SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
SET GIN_MODE=release
go build -o workbench.exe .