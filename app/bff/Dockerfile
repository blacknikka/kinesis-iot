FROM golang:1.14-alpine as build

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/blacknikka/kinesis-iot
WORKDIR /go/src/github.com/blacknikka/kinesis-iot

COPY ./src /go/src/github.com/blacknikka/kinesis-iot
RUN go get
RUN go build -o main

FROM alpine:3.12

COPY --from=build /go/src/github.com/blacknikka/kinesis-iot/main /app/main
WORKDIR /app
RUN chmod +x /app/main

CMD [ "/app/main" ]
