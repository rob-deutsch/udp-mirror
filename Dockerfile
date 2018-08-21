FROM golang

RUN apt-get install -y git

RUN go get github.com/rob-deutsch/udp-mirror

ENTRYPOINT [ "udp-mirror" ]
EXPOSE 9999
