FROM debian:stable-slim

COPY troll /usr/local/bin/troll

ENV PORT=8080

EXPOSE $PORT

ENTRYPOINT ["/usr/local/bin/troll"]