package main

//this go file contains calculations functions and validating functions of the input by the client

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

// providing a simple structure to perform simple math operration
type Operation struct {
	num1, num2 float64
	operator   string
}

//testing the input for the right ampunt of parameters
func testInputString(input []string) error {
	if len(input) != 3 {
		return status.Error(codes.InvalidArgument, "Wrong amount of parameters in the input")
	}
	return nil
}

// casting the numbers from string format to float for further calculations
func parseArgs(c []string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return 0.0, 0.0, status.Error(codes.InvalidArgument, "Invalid first argument")
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		return 0.0, 0.0, status.Error(codes.InvalidArgument, "Invalid second argument")
	}
	return num1, num2, nil
}

// mapping the operator to an allowed math operation and returning the result
func processOperation(op Operation) (float64, error) {
	var result float64
	switch op.operator {
	case "*":
		result = op.num1 * op.num2
	case "/":
		if op.num2 == 0.0 {
			return 0.0, status.Error(codes.InvalidArgument, "error: you tried to divide by zero.")
		}
		result = op.num1 / op.num2
	case "+":
		result = op.num1 + op.num2
	case "-":
		result = op.num1 - op.num2
	default:
		return 0.0, status.Error(codes.InvalidArgument, "Not acceptable Operation")
	}
	return result, nil
}
