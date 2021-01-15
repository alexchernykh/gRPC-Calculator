FROM golang:1.15

ENV GO111MODULE=on

WORKDIR /go/src/gRPC-Calculator
COPY ./client/* ./client/
COPY ./go.mod ./
COPY ./go.sum ./
COPY ./calc_pb/* ./calc_pb/

#EXPOSE 8080

RUN go mod download

CMD ["bash","client/client.sh"]