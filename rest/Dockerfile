FROM golang:1.16-alpine

WORKDIR /usr/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /sky-api

EXPOSE 8081

CMD [ "/sky-api" ]