build-cli:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "-N -l"  -o bin/amd64/asm-package main.go
build-cli-arm:
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -gcflags "-N -l"  -o bin/arm64/asm-package main.go
