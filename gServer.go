package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
	p "testAssignmment/calc_pb"
)
var port = ":8080"

//required in order to be able to create the gRPC server later on in your Go code.
type MessageServer struct {

}

// providing a simple structure to perform simple math operration
type Operation struct {
	num1, num2 float64
	operator   string
}

//testing the input for the right ampunt of parameters
func testInputString(input []string) (error)  {
	if len(input) !=3 {
		return errors.New("Wrong amount of parameters in the input")
	}else{
		return  nil
	}
}

// casting the numbers from string format to float for further calculations
func parseArgs(c []string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	return num1, num2, nil
}

// mapping the operator to an allowed math operation and returning the result
func processOperation(op Operation) (float64, error){
	var result float64
	switch op.operator {
	case "*":
		result = op.num1 * op.num2
	case "/":
		if op.num2 == 0.0 {
			return 0.0, errors.New("error: you tried to divide by zero.")
		}
		result = op.num1 / op.num2
	case "+":
		result = op.num1 + op.num2
	case "-":
		result = op.num1 - op.num2
	default:
		return 0 , errors.New("Not acceptable Operation: ")
	}
	return result, nil
}

// this is the implementation of the interface from calc.pd.go, that will be triggered by the client
// Takes a request from the cluent and checks the input for errors , performs the simple math
// operation and returns the result to the client
func (MessageServer) CalcResult(ctx context.Context, r *p.Request) (*p.Response, error){
		fmt.Println(r.Subtext , "for ", r.Text)
		c := strings.Split(r.Text, " ")
		err := testInputString(c)
		if err != nil{
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
		} else {
			res_str := fmt.Sprintf("%f", res)
			response := &p.Response{
				Text:res_str,
				Subtext:"Successfully finished calculation!",
			}
			return response, nil
		}
}


func main() {

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