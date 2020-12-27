package main

import (
	"fmt"
	p "github.com/alexchernykh/gRPC-Calculator/calc_pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"net"
	"strings"
)

//required in order to be able to create the gRPC server later on in your Go code.
type MessageServer struct {
}

// this is the implementation of the interface from calc.pd.go, that will be triggered by the client
// Takes a request from the cluent and checks the input for errors , performs the simple math
// operation and returns the result to the client
func (MessageServer) CalcResult(ctx context.Context, r *p.Request) (*p.Response, error) {
	fmt.Println(r.Subtext, "for ", r.Text)
	c := strings.Split(r.Text, " ")
	err := testInputString(c)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var op = Operation{operator: c[1]}

	//parse the numbers from the input string
	op.num1, op.num2, err = parseArgs(c)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//perform the Operation on the two numbers
	res, err := processOperation(op)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res_str := fmt.Sprintf("%f", res)
	response := &p.Response{
		Text:    res_str,
		Subtext: codes.Code.String(codes.OK),
	}
	fmt.Println("Successfully finished calculation!")
	return response, nil

}

func main() {
	var port = ":8080"
	server := grpc.NewServer()
	var messageServer MessageServer

	p.RegisterMessageServiceServer(server, messageServer)
	listen, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println(err)
		return
	}
	//listen to the port for new remote procedure calls
	fmt.Println("Starting serving requests...")
	server.Serve(listen)

}
