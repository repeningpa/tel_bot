FROM golang:1.12.0-alpine3.9

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN apt install -y git
RUN go get github.com/Syfaro/telegram-bot-api
RUN go build -o main .

CMD ["/app/main"]