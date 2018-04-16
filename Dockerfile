FROM resin/rpi-raspbian:latest

ENTRYPOINT []

RUN apt-get update && \
    apt-get -qy install curl ca-certificates

ADD translatebot translatebot
RUN chmod +x translatebot

CMD ["./translatebot"]
