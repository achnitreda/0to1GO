FROM golang:1.22.3-alpine AS start

WORKDIR /ascii-art-web-dockerize

COPY . .

WORKDIR /ascii-art-web-dockerize/server

RUN  go build -o web

FROM alpine

LABEL version="1.0" description="Dockerized ASCII Art Web Application"

RUN apk add bash

WORKDIR /root

COPY --from=start /ascii-art-web-dockerize .

WORKDIR /root/server

EXPOSE 8081

CMD [ "./web" ]