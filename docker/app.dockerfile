#FROM golang:1.20.4-alpine3.18
FROM golang:1.20.4-bullseye
#ARG UID=1000
#ARG GID=1000
#RUN groupadd -g 1000 -o appuser
#RUN useradd -r -u $UID -g $GID appuser
#USER appuser

WORKDIR /goAPIrest

ADD . .

RUN go mod download

ENTRYPOINT go build -o apirest main.go && ./apirest