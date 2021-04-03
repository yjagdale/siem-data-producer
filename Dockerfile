FROM centos:8

WORKDIR /home
ENV GIN_MODE=release
ADD siem-data-producer siem-data-producer
ADD static static
ENTRYPOINT ["/home/siem-data-producer"]