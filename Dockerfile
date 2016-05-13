FROM golang:latest

WORKDIR /app
COPY app /app

EXPOSE 3000
