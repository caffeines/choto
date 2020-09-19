FROM golang:alpine AS builder

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add git openssh

ENV GOPATH=/go

ENV GOOS="linux"
ENV GOARCH="amd64"
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/caffeines/choto
COPY . $GOPATH/src/github.com/caffeines/choto

RUN go mod download
RUN go build -v -o choto
RUN mv choto /go/bin/choto

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /root

COPY --from=builder /go/bin/choto /usr/local/bin/choto

EXPOSE 4521

CMD ["choto"]