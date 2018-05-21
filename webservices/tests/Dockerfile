FROM golang:1.9 AS build

WORKDIR /go/tests
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -a -o /users users.go
FROM alpine
COPY --from=build /users ./users
ENTRYPOINT ./tests