package main

import (
	"bufio"
	"fmt"
	p "github.com/alexchernykh/gRPC-Calculator/calc_pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	"strings"
	"flag"
)

//Takes a string from user input return it without the new line suffix
func readInputData() string {
	fmt.Println("Please Enter an Operation that has to be done in a form (a + b), with supported operations *,/,+,- :")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}



//Takes an input string from STDIN, preprocesses it and sends it to the server by using the
//MessageServiceClient, that is a gRPC ClientConnInterface implemented in the "testAssignmment/calc_pb"
func ReadyToCalc(ctx context.Context, m p.MessageServiceClient) (*p.Response, error) {
	input := readInputData()
	fmt.Println("This was the input value: ", input)

	//create request structure to send to CalcResult
	request := &p.Request{
		Text:    input,
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

	//port based on which client will create a connection.
	// to speak to a localhost there is no need to provide an container name like server.
	// However, when using docker container for client and separate for server, both containers will have separate localhosts
	// and require a name to be provided.
	//var port = "server:8080"
	port := flag.String("port", ":8080", "a string")
	flag.Parse()

	conn, err := grpc.Dial(*port, grpc.WithInsecure())
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
