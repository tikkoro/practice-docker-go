FROM golang:1.19.1

ARG USER
ARG PASSWORD
ARG NAME

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/go/app/src
ENV DB_USER=${USER}
ENV DB_PASS=${PASSWORD}
ENV DB_NAME=${NAME}

RUN apt update && apt install -y nano
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/cosmtrek/air@latest
WORKDIR ${ROOTPATH}
