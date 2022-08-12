FROM golang:alpine AS builder

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add git openssh

ENV GOPATH=/go

ENV GOOS="linux"
ENV GOARCH="amd64"
ENV GO11MODULE="on"

COPY . $GOPATH/src/github.com/tawsifkarim/do-app
WORKDIR $GOPATH/src/github.com/tawsifkarim/do-app

RUN go get .

RUN go build -v -o do-app
RUN mv tracking-service /go/bin/do-app

### Step 2

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /root

COPY --from=builder /go/bin/do-app /usr/local/bin/do-app

# ENTRYPOINT ["do-app"]
CMD [ "do-app" ]