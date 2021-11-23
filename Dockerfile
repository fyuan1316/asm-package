ARG image_registry=harbor-b.alauda.cn/asm
ARG runner_base_tag=3.14.0
FROM golang:1.16-alpine AS builder
WORKDIR /workspace
ENV GOPROXY=https://goproxy.cn,direct

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# src code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "-N -l"  -o bin/amd64/asm-package main.go && \
     CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -gcflags "-N -l"  -o bin/arm64/asm-package main.go


FROM alpine:3.7
COPY --from=builder /workspace/bin /asm-package/bin
COPY artifacts /asm-package

ENV TZ=Asia/Shanghai




