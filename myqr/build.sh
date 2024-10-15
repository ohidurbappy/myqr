# go build -ldflags="-s -w -H=windowsgui" -o 
GOOS=darwin GOARCH=arm64 go build -o "../build/myqr"
GOOS=windows GOARCH=amd64 go build -o "../build/myqr.exe"
