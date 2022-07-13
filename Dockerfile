# Dockerfile多阶构建
# Pull the base image
FROM golang:1.18 as builder
# Setup the maintainer.
MAINTAINER hurricane19898@faw.cn
# Setup labels.
LABEL image.authors="hurricane"
LABEL image.documentation="https://ymmt2005.hatenablog.com/entry/2020/04/14/An_example_of_using_dynamic_client_of_k8s.io/client-go"

# Add all the files to contaier
ADD . /build/
# Setup work directory.
WORKDIR /build

# Setup the go proxy.
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on
# 设置编译环境
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -ldflags '-s -w' -installsuffix cgo -o main

# Builde the final image.
FROM scratch
# Setup the healthcheck.
# HEALTHCHECK --interval=5s --timeout=3s --retries=3 CMD curl -fs http://127.0.0.1:8080/ || exit 1
COPY --from=builder /build/main /
# Expose the tcp 8080 port.
EXPOSE 8080
ENTRYPOINT ["/main"]