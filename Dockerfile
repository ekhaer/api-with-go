FROM golang:1.18-alpine

LABEL maintainer="@ekhaer"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

RUN go get -d -v ./...

RUN go install -v ./...

RUN source .env

RUN go build -o /build

EXPOSE 8080

CMD [ "/build" ]