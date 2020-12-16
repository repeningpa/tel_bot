FROM golang:1.9.2 

ADD . /go/src/app

WORKDIR /go/src/app

RUN go get main

RUN go install

ENTRYPOINT ["/go/bin/main"]
