# dependencies
# RUN go get labix.org/v2/mgo
# RUN go get github.com/julienschmidt/httprouter

FROM golang:latest

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app

RUN go-wrapper download
RUN go-wrapper install

ENV PORT 3000
EXPOSE 3000

CMD ["go-wrapper", "run"]
