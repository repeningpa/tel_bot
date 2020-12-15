FROM golang:latest

COPY ./simpleWebServer/ ./build
RUN go build -o main .
CMD ["./simpleWebServer/main"]
