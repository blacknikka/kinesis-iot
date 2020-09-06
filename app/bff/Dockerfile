FROM golang:1.14-alpine

RUN mkdir -p /go/src/github.com/blacknikka/kinesis-iot
WORKDIR /go/src/github.com/blacknikka/kinesis-iot

COPY ./src /go/src/github.com/blacknikka/kinesis-iot
RUN go get

CMD [ "/bin/ash" ]
