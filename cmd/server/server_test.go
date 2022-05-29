package main

import (
	"errors"
	"fmt"
	"testing"

	calc "github.com/zahaar/grcp-calculator/gen"
)

func TestCalculate(t *testing.T) {
	// test that calculate() calc.MathMethod_ADD works
	req := &calc.MathTaskRequest{Method: calc.MathMethod_ADD, Arg1: 21, Arg2: 21}
	actual, _ := calculate(req)
	var expected float64 = 42
	if actual.GetEq() != expected {
		t.Error("Should return 42")
	}

	// test calculate() addition of negative values
	req = &calc.MathTaskRequest{Method: calc.MathMethod_ADD, Arg1: -17, Arg2: -25}
	actual, _ = calculate(req)
	expected = -42
	if actual.GetEq() != expected {
		t.Error("Should return -42")
	}

	// test calculate() addition of float values
	req = &calc.MathTaskRequest{Method: calc.MathMethod_ADD, Arg1: 17.17, Arg2: 24.83}
	actual, _ = calculate(req)
	expected = 42
	if actual.GetEq() != expected {
		t.Error("Should return -42")
	}

	// test calculate() division of float values
	req = &calc.MathTaskRequest{Method: calc.MathMethod_DIV, Arg1: 17.17, Arg2: -24.5}
	actual, _ = calculate(req)
	expected = -0.7008163265306123
	if actual.GetEq() != expected {
		fmt.Println(actual)
		t.Error("Should return -7.659999999999997")
	}

}

func TestCheckMethodExistDivideByZero(t *testing.T) {
	// test calculate() division by zero returns err
	req := &calc.MathTaskRequest{Method: calc.MathMethod_DIV, Arg1: 17.17, Arg2: 0.0}
	_, err := calculate(req)
	if errors.Is(err, errors.New("math: divided by zero")) {
		t.Error("Should return an err")
	}
}
