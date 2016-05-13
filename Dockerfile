FROM golang:latest

ADD ./app /go/src/app/

RUN go install app/

ENTRYPOINT /go/bin/app

EXPOSE 3000
