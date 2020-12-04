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

//Takes a string from user input return it without the new line suffix
func readInputData() string {
	fmt.Println("Please Enter an Operation that has to be done in a form (a + b), with supported operations *,/,+,- :")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}


//port based on which client will create a connection.
var port = ":8080"


//Takes an input string from STDIN, preprocesses it and sends it to the server by using the
//MessageServiceClient, that is a gRPC ClientConnInterface implemented in the "testAssignmment/calc_pb"
func ReadyToCalc(ctx context.Context, m p.MessageServiceClient) (*p.Response, error) {
	input := readInputData()
	fmt.Println("This was the input value: ",input)

	//create request structure to send to CalcResult
	request := &p.Request{
		Text:input,
		Subtext: "New Calculation",
	}
	//call the remote func via the messageServiceClient
	calculated_result, err := m.CalcResult(ctx, request)
	if err != nil {
		return nil, err
	}
	return calculated_result, nil
}

func main() {

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}
	client := p.NewMessageServiceClient(conn)

	//
	r, err := ReadyToCalc(context.Background(), client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response Text:", r.Text)
	fmt.Println("Response SubText:", r.Subtext)

}