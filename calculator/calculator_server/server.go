package main

import (
	"context"
	"fmt"
	//"io"
	"log"
	"net"
	//"strconv"
	//"time"

	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"

	"github.com/suganoo/go-grpc-course/calculator/calculatorpb"

	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Received Sum RPC: %v\n", req)
	firstNumber  := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
