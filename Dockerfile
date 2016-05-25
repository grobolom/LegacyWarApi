FROM golang:latest

RUN mkdir -p app/
WORKDIR app/

COPY . app/

# dependencies
# RUN go get labix.org/v2/mgo
RUN go get github.com/julienschmidt/httprouter

EXPOSE 3000

CMD ["go", "run", "main.go"]
