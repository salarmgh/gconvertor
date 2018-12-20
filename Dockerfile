FROM golang:1.11.3-alpine AS builder

ARG HTTP_PROXY
ARG HTTPS_PROXY

RUN apk add --update git &&  wget https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -O /usr/local/bin/dep && chmod +x /usr/local/bin/dep

WORKDIR $GOPATH/src/github.com/salarmgh/gconvertor
COPY . ./
RUN HTTP_PROXY=$HTTP_PROXY HTTPS_PROXY=$HTTPS_PROXY dep ensure && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM alpine:3.8 AS app
RUN apk add --update --no-cache ffmpeg
COPY --from=builder /app ./
WORKDIR "/data"
ENTRYPOINT ["/app"]
