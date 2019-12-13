FROM golang:1.12-alpine AS build-env
COPY . /go/src/app
WORKDIR /go/src/app/websocket/q1
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.io go build -v -a -installsuffix cgo -o app-server

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/app/websocket/q1/app-server .
ENTRYPOINT ["/app/app-server"]
EXPOSE 2303