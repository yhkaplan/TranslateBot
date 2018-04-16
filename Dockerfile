FROM resin/rpi-raspbian:latest

ENTRYPOINT []

ADD translatebot translatebot
RUN chmod +x translatebot

CMD ["./translatebot"]
