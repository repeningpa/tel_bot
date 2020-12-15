FROM golang:latest

COPY ./../simpleWebServer/main ./build
RUN go build -o main .
CMD ["./../simpleWebServer/main"]
