FROM golang:latest

WORKDIR /go_api/src

RUN apt-get update \
&& apt-get install -y git

CMD ["go","run","main.go"]