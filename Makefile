# Makefile

build_linux:
	mkdir -p ./output
	CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -a -ldflags="-w -s" -o ./output/githubhost_linux && upx -9 ./output/githubhost_linux

build_windows:
	mkdir -p ./output
	CGO_ENABLED=0 GOOS=windows  GOARCH=amd64 go build -a -ldflags="-w -s" -o ./output/githubhost_windows.exe && upx -9 ./output/githubhost_windows.exe

build_darwin:
	mkdir -p ./output
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -a -ldflags="-w -s" -o ./output/githubhost_darwin && upx -9 ./output/githubhost_darwin

build_all: build_linux build_darwin build_windows


build:
	mkdir -p ./output
	go build -a -ldflags="-w -s" -o ./output/githubhost_local
	chmod +x ./output/githubhost_local

run: build
	sudo ./output/githubhost_local








