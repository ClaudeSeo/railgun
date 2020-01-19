FROM golang:1.13.6-alpine AS builder

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build

FROM alpine:latest

RUN apk update

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY --from=builder /usr/src/app .

EXPOSE 8080

CMD [ "./railgun" ]
