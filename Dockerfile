FROM alpine
WORKDIR /app
ADD main main
EXPOSE 8080
ENTRYPOINT ["/app/main"]