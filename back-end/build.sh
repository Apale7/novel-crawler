GOOS=linux GOARCH=amd64 go build -o server_linux.out
GOOS=windows GOARCH=amd64 go build -o server_windows.exe
GOOS=darwin GOARCH=amd64 go build -o server_osx.out