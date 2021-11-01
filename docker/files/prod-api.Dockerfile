FROM golang:alpine AS build

# GOPROXY resolves dependencies treefrom cache or repository
ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/feeder-service
COPY . .
# Set OS as linux
RUN GOOS=linux go build -o /go/bin/feeder-service cmd/api/main.go

FROM alpine
COPY --from=build /go/bin/feeder-service /go/bin/feeder-service
ENTRYPOINT ["go/bin/feeder-service"]
