FROM golang:1.15 AS build

ENV GO111MODULE=on

WORKDIR /go/src/gRPC-Calculator
COPY ./server/* ./server/
COPY ./go.mod ./
COPY ./go.sum ./
COPY ./calc_pb/* ./calc_pb/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server/server ./server/*.go

FROM scratch
WORKDIR /go/src/gRPC-Calculator
COPY --from=build /go/src/gRPC-Calculator/server/server /go/src/gRPC-Calculator
EXPOSE 8080
CMD ["./server"]