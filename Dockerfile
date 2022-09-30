## Build
FROM golang:1.19-buster AS build


COPY . /usr/src/troll/

WORKDIR /usr/src/troll

RUN go build -o /usr/local/bin/troll


## Deploy
FROM debian:stable-slim

COPY --from=build /usr/local/bin/troll /usr/local/bin/troll
COPY templates /opt/troll/templates
COPY static /opt/troll/static
COPY public /opt/troll/public
COPY v2_api.yaml /opt/troll/v2_api.yaml


WORKDIR /opt/troll

ENV ADDRESS=":8080"

EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/troll"]