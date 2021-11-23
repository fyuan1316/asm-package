build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "-N -l"  -o bin/amd64/asm-package main.go
build-linux-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -gcflags "-N -l"  -o bin/arm64/asm-package main.go
build:
	CGO_ENABLED=0 GOARCH=amd64 go build -gcflags "-N -l"  -o bin/asm-package main.go
clean:
	rm -rf bin
docker-build:
	docker buildx use mbuilder
	docker buildx build --platform linux/amd64,linux/arm64 -t build-harbor.alauda.cn/asm/package-tools --push .