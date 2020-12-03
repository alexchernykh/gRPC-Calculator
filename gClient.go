package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	"strings"
	p "testAssignmment/calc_pb"
)

//Takes a string from user input and
func readInputData() string {
	fmt.Println("Please Enter an operation that has to be done in a form (a + b), with supported operations *,/,+,- :")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	//c := strings.Split(input, " ")
	//joined := strings.Join(c," ")
	return input
}

var port = ":8080"
//Takes an input string from STDIN, preprocesses it and sends it to the server by using the
//MessageServiceClient, that is a gRPC ClientConnInterface implemented in the "testAssignmment/calc_pb"
func ReadyToCalc(ctx context.Context, m p.MessageServiceClient) (*p.Response, error) {
	input :=readInputData()
	fmt.Println("This was the input value: ",input)
	request := &p.Request{
		Text:input,
		Subtext: "New Calculation!",
	}
	//вызов функции с сервера
	r, err := m.CalcResult(ctx, request)
	//r, err := m.SayIt(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}
	client := p.NewMessageServiceClient(conn)
	r, err := ReadyToCalc(context.Background(), client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response Text:", r.Text)
	fmt.Println("Response SubText:", r.Subtext)
}