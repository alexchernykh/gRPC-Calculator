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
type MessageServer struct {

}


type operation struct {
	num1, num2 float64
	operat string
}

func parseArgs(c []string) (float64, float64, error) {
	if len(c)<3{
		return 0.0, 0.0, errors.New("Not enough arguments")
	}
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


func processOperation(op operation) (float64, error){
	var result float64
	switch op.operat {
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
		return 0 , errors.New("Not acceptable operation: ")
	}
	return result, nil
}

func (MessageServer) CalcResult(ctx context.Context, r *p.Request) (*p.Response, error){


		//fmt.Println("Please Enter an operation that has to be done in a form (a + b ), with supported operations *,/,+,- :")
		//reader := bufio.NewReader(os.Stdin)
		//input, _ := reader.ReadString('\n')
		//input = strings.TrimSuffix(input, "\n")
		c := strings.Split(r.Text, " ")
		var op = operation{operat: c[1]}
		var err error
		op.num1, op.num2, err = parseArgs(c)
		if err != nil {
			fmt.Println(err)
			response := &p.Response{

				Text:"Error",

				Subtext:"Error",
			}
			return response, err
		}
		//res, err := processOperation(num1, c[1], num2)
		res, err := processOperation(op)
		if err != nil {
			fmt.Println(err)
			// todo return nothing in response and error
			fmt.Println(err)
			response := &p.Response{

				Text:"Error",

				Subtext:"Error",
			}
			return response, err
		} else {
			res_str := fmt.Sprintf("%f", res)
			response := &p.Response{

				Text:res_str,

				Subtext:"Got it!",
			}
			return response, nil
			//fmt.Println("Calculated result  is = ", res)
		}

}




//
//func (MessageServer) SayIt(ctx context.Context, r *p.Request) (*p.Response, error) {
//
//	fmt.Println("Request Text:", r.Text)
//
//	fmt.Println("Request SubText:", r.Subtext)
//
//	response := &p.Response{
//
//		Text:r.Text,
//
//		Subtext:"Got it!",
//	}
//	return response, nil
//}

func main() {

	server := grpc.NewServer()

	var messageServer MessageServer

	p.RegisterMessageServiceServer(server, messageServer)

	listen, err := net.Listen("tcp", port)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("Serving requests...")

	server.Serve(listen)

}