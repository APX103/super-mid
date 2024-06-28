FROM golang:1.21.0-alpine3.18 as builder
WORKDIR /workspace/
ADD ./ /workspace/

ENV GO_PROXY="https://goproxy.cn,direct"

RUN go env -w GOPROXY=${GO_PROXY}
RUN sh ./scripts/build.sh

FROM alpine:3.18.2

COPY --from=builder /workspace/super-mid /workspace/super-mid
COPY ./config /workspace/config
WORKDIR /workspace

EXPOSE 8080

RUN chmod +x super-mid

# set timezone
RUN sed -i 's@dl-cdn.alpinelinux.org@mirrors.aliyun.com@g' /etc/apk/repositories
RUN apk add -U tzdata && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && apk del tzdata

CMD ./super-mid 2>&1