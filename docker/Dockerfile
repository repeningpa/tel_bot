FROM golang:1.9.2 

ADD . /go/src/main

WORKDIR /go/src/main

RUN go get main

RUN go install

ENTRYPOINT ["/go/bin/main"]
