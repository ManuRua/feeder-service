FROM golang:stretch AS build-debug

EXPOSE 40000

WORKDIR /go/src/feeder-service

RUN go get -u github.com/go-delve/delve/cmd/dlv

# copying all from context .. to /go/src/feeder-service
COPY . .

RUN go mod download

ENTRYPOINT dlv debug ./cmd --listen=:40000 --headless=true --api-version=2
