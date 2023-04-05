FROM golang:1.19 as builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .    

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

RUN mkdir publish && \
    cp goblog publish && \
    cp adm.ini publish && \
    cp config.yml.example publish/config.yml && \
    cp redis.yml.example publish/redis.yml && \
    cp sentry.yml.example publish/sentry.yml && \
    cp rocketmq.yml.example publish/rocketmq.yml && \
    mkdir logs

FROM alpine:3.14

RUN apk add wait4x

WORKDIR /app

COPY --from=builder /app/publish .

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/cert

ENV GIN_MODE=release \
    PORT=80

EXPOSE 80    

# ENTRYPOINT [ "./goblog" ]