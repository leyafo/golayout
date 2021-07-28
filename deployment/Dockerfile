FROM golang:1.16-alpine AS builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories &&\
    apk update &&\
    apk --no-cache add make git build-base linux-headers

WORKDIR /opt/ops
COPY . .

#build at China
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    make clean && make build

FROM alpine:latest

WORKDIR /opt/ops

COPY deployment/entrypoint.sh /
COPY deployment/glibc-2.33-r0.apk /tmp/
COPY --from=builder /opt/ops/bin/ /opt/ops/bin/
COPY --from=builder /opt/ops/conf/ /opt/ops/conf/

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g" /etc/apk/repositories &&\
   apk update && apk --no-cache add tini tzdata ca-certificates && \
   touch repo.list && apk add --repositories-file=repo.list --allow-untrusted --no-network --no-cache /tmp/glibc-2.33-r0.apk &&\
   mkdir -p /opt/ops/log && \
   cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
   chmod +x /entrypoint.sh /opt/ops/bin/*

ENV PATH /opt/ops/bin:$PATH

ENTRYPOINT ["/sbin/tini", "--", "/entrypoint.sh"]