FROM golang as build


ADD . /usr/local/go/src/blog

WORKDIR /usr/local/go/src/blog

RUN GOPROXY="https://goproxy.cn" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_blog

FROM alpine:3.7
#
#ENV GIN_MODE="release"
#ENV PORT=3000

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /usr/local/go/src/blog/api_blog /usr/bin/api_blog

ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_blog

ENTRYPOINT ["api_blog"]
