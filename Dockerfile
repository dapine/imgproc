FROM golang:alpine
MAINTAINER Breno Pinheiro <davepinh@gmail.com>
WORKDIR /app

RUN apk update && apk upgrade
RUN apk add musl-dev vips-dev gcc

COPY . .
RUN go mod download
RUN go build
CMD ["./imgproc"]
EXPOSE 9001
