FROM golang:alpine as dlv-env

RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM alpine:latest

WORKDIR /app

COPY --from=dlv-env /go/bin/dlv /go/bin/dlv

ADD main main

EXPOSE 8000 40000

#ENTRYPOINT ["/app/main"] # run without debugger
#ENTRYPOINT ["/go/bin/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "--log", "exec", "/app/main"] # run with debugger, wait for debugger to attach
#ENTRYPOINT ["/go/bin/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "--log", "--continue", "exec", "/app/main"] # run with debugger, do not wait for debugger to attach
