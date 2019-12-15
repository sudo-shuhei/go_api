FROM golang:latest

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . /go/src/work

RUN go get -u github.com/labstack/echo/...
RUN go get github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
