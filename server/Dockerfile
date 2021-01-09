FROM golang:1.15

ENV GO111MODULE=on

WORKDIR /go/src/gRPC-Calculator
COPY ./server/* ./server/
COPY ./go.mod ./
COPY ./go.sum ./
COPY ./calc_pb/* ./calc_pb/

EXPOSE 8080

RUN go mod download

CMD ["go","run","server/calcuations.go","server/gServer.go"]