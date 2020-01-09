FROM golang:latest

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . /go/src/work

RUN go get -u github.com/labstack/echo/...
RUN go get github.com/lib/pq
RUN go get -u github.com/jinzhu/gorm
RUN go get github.com/tools/godep
