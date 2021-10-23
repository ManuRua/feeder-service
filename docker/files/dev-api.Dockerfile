FROM golang:stretch AS build-dev

# GOPROXY resolves dependencies treefrom cache or repository
ENV GOPROXY=https://proxy.golang.org
# ENV GO111MODULE=on
WORKDIR /go/src/feeder-service
# copying all from context .. to /go/src/feeder-service
COPY . .

RUN go mod download 
RUN go get -u github.com/cosmtrek/air

ENTRYPOINT air -d