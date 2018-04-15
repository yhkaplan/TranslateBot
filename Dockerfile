FROM resin/rpi-raspbian:latest

ENTRYPOINT []

RUN apt-get update && \
    apt-get -qy install curl ca-certificates

ADD TranslateBot TranslateBot
RUN chmod +x TranslateBot

CMD ["./TranslateBot"]
