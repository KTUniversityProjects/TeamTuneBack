FROM golang:1.9.2-alpine3.6

COPY . /go/
RUN export GOPATH=$HOME/project/
RUN go get gopkg.in/mgo.v2
RUN go get golang.org/x/crypto/bcrypt
WORKDIR /go/src/core
RUN go install
EXPOSE 80
