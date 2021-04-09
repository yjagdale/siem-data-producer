FROM golang:1.16.3

WORKDIR /home
ENV GIN_MODE=release
ADD siem-data-producer siem-data-producer
ADD static static
ENTRYPOINT ["/home/siem-data-producer"]