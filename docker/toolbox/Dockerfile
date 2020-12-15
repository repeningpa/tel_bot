FROM golang:latest

COPY ./ ./build
RUN go build -o main .
CMD ["./../simpleWebServer/main"]
