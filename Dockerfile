
FROM golang:onbuild
RUN apt-get upgrade -y

RUN go get encoding/json
RUN  go get -u github.com/gorilla/mux
RUN go get database/sql

ENV GOBIN /go/bin
RUN go get github.com/go-sql-driver/mysql
EXPOSE 8080
MAINTAINER email
