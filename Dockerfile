FROM golang:1.11.3 AS builder

ARG HTTP_PROXY
ARG HTTPS_PROXY

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/salarmgh/gconvertor
COPY . ./
RUN HTTP_PROXY=$HTTP_PROXY HTTPS_PROXY=$HTTPS_PROXY dep ensure && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch AS app
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
