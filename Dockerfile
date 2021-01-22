FROM ubuntu:18.04

WORKDIR /home
ADD siem-data-producer siem-data-producer
ADD static static
ENTRYPOINT ["/home/siem-data-producer"]