FROM golang:alpine AS build
WORKDIR /usr/gateway
COPY socket/. .
RUN go get
RUN apk add gcc musl-dev libc-dev make && \
    GOZSTD_VER=$(cat go.mod | fgrep github.com/valyala/gozstd | awk '{print $NF}') && \
    go get -d github.com/valyala/gozstd@${GOZSTD_VER} && \
    cd ${GOPATH}/pkg/mod/github.com/valyala/gozstd@${GOZSTD_VER} && \
    make clean && \
    make -j $(nproc) libzstd.a && \
    cd /usr/gateway && \
    go build main.go

FROM alpine:latest
COPY --from=build /usr/gateway/main /gateway
ENTRYPOINT ["/gateway"]