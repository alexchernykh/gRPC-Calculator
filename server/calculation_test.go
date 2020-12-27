package main

import (
	"log"
	"strings"
	"testing"
)

func TestTestInputString(t *testing.T) {
	equation := "2 + 3"
	c := strings.Split(equation, " ")
	result := testInputString(c)
	if result != nil {
		t.Errorf("testInputString() failed, expected nil, got  %v", result)
	}
}

func TestParseArgs(t *testing.T) {
	equation := "2 + 3"
	c := strings.Split(equation, " ")
	var op Operation
	var err error
	op.num1, op.num2, err = parseArgs(c)
	if err != nil {
		log.Panic("Error occured in parseArgs()")
	}
	if op.num1 != 2 {
		t.Errorf("parseArgs() failed for first argument, expected 2, got  %v", op.num1)
	}

	if op.num2 != 3 {
		t.Errorf("parseArgs() failed for third argument, expected 3, got  %v", op.num1)
	}
}

func TestProcessOperationPlus(t *testing.T) {
	var op = Operation{1, 2, "+"}
	var err error
	var result float64
	result, err = processOperation(op)
	if err != nil {
		log.Panic("Error occured in processOperation()")
	}
	if result != 3 {
		t.Errorf("processOperation() failed, expected 3, got  %v", result)
	}
}

func TestProcessOperationSubstract(t *testing.T) {
	var op = Operation{2, 1, "-"}
	var err error
	var result float64
	result, err = processOperation(op)
	if err != nil {
		log.Panic("Error occured in processOperation()")
	}
	if result != 1 {
		t.Errorf("processOperation() failed, expected 1, got  %v", result)
	}
}

func TestProcessOperationMultiplication(t *testing.T) {
	var op = Operation{3, 4, "*"}
	var err error
	var result float64
	result, err = processOperation(op)
	if err != nil {
		log.Panic("Error occured in processOperation()")
	}
	if result != 12 {
		t.Errorf("processOperation() failed, expected 12, got  %v", result)
	}
}

func TestProcessOperationDivision(t *testing.T) {
	var op = Operation{12, 4, "/"}
	var err error
	var result float64
	result, err = processOperation(op)
	if err != nil {
		log.Panic("Error occured in processOperation()")
	}
	if result != 3 {
		t.Errorf("processOperation() failed, expected 3, got  %v", result)
	}
}
