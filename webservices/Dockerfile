FROM golang:1.9.2-alpine3.6 AS build

# Install tools required to build the project
# We need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git
RUN go get gopkg.in/mgo.v2
RUN go get golang.org/x/crypto/bcrypt

COPY . /go/src/project/
export GOPATH=$HOME/go/src/project/core/structures
export GOPATH=$HOME/go/src/project/core

WORKDIR /go/src/project/services/boards
RUN go build -o /bin/project

FROM scratch
COPY --from=build /bin/project /bin/project
ENTRYPOINT ["/bin/project"]
CMD ["--help"]

