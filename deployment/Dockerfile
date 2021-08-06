FROM golang:1.16-alpine AS builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories &&\
    apk update &&\
    apk --no-cache add make git build-base linux-headers

WORKDIR /opt/app
COPY . .

#build at China
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    make clean && make build

FROM alpine:latest

WORKDIR /opt/app

COPY deployment/entrypoint.sh /
COPY --from=builder /opt/app/bin/ /opt/app/bin/
COPY --from=builder /opt/app/conf/ /opt/app/conf/

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g" /etc/apk/repositories &&\
   apk update && apk --no-cache add tini tzdata ca-certificates && \
   mkdir -p /opt/app/log && \
   cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
   chmod +x /entrypoint.sh /opt/app/bin/*

ENV PATH /opt/app/bin:$PATH

ENTRYPOINT ["/sbin/tini", "--", "/entrypoint.sh"]