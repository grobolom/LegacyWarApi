FROM golang:latest

ADD ./app /go/src/app/

# dependencies
RUN go get labix.org/v2/mgo
RUN go get labix.org/v2/bson

RUN go install app/

ENTRYPOINT /go/bin/app

EXPOSE 3000
