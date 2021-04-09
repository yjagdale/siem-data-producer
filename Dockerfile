FROM debian:9.5-slim

WORKDIR /home
RUN apt-get install libc6-dev=2.17-7 -y
ENV GIN_MODE=release
ADD siem-data-producer siem-data-producer
ADD static static
ENTRYPOINT ["/home/siem-data-producer"]