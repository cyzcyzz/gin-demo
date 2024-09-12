# Build step11
FROM golang:1.18.10 AS builder
RUN  mkdir -p  /home/work/data/www/cyztest
WORKDIR /home/work/data/www/cyztest
COPY . .
RUN ls ./
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.io,direct && go build -o /cyztest
# Final step
FROM alpine:3.17
MAINTAINER chengyuanzhao "chengyuanzhao@xiaomi.com"
WORKDIR /home/work/data/www/cyztest
COPY --from=builder /cyztest .
