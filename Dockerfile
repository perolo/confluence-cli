FROM golang:1.14-alpine as builder
ADD . /go/src/github.com/proctorlabs/confluence-cli/
RUN apk add git
RUN go get golang.org/x/net/html
RUN cd /go/src/github.com/proctorlabs/confluence-cli/ && go build main.go

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/proctorlabs/confluence-cli/main /usr/local/bin/confluence-cli
ENTRYPOINT [ "confluence-cli" ]
