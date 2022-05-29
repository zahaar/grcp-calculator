package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	calc "github.com/zahaar/grcp-calculator/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	method = flag.String("method", "", "Math methods: add, sub, mul, div")
	arg1   = flag.Float64("a", 0, "Math method arg1")
	arg2   = flag.Float64("b", 0, "Math method arg2")
)

func checkMethodExist(m *string) calc.MathMethod {
	val, pres := calc.MathMethod_value[strings.ToUpper(*m)]

	if !pres {
		log.Fatalf("Method: {%v}, is not present in: add, sub, mul, div", *m)
	}
	return calc.MathMethod(val)
}

func main() {
	flag.Parse()
	calcMethod := checkMethodExist(method)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	c := calc.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.PerformCalc(ctx, &calc.MathTaskRequest{Method: calcMethod, Arg1: float64(*arg1), Arg2: float64(*arg2)})
	if err != nil {
		log.Fatalf("fail to perform MathTaskRequest() err:%v", err)
	}
	log.Printf("Result: %v", res.GetEq())
}
