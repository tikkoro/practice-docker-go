FROM golang:1.19.1

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/go/app/src

RUN apt update
RUN apt install -y nano

WORKDIR ${ROOTPATH}
