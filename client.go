package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("Error connecting to RPC server: %v", err)
	}
	defer client.Close()

	args := Args{A: 10, B: 5}
	var reply int
	err = client.Call("Calculator.Add", args, &reply)
	if err != nil {
		log.Fatalf("Error calling Calculator.Add: %v", err)
	}
	fmt.Printf("10 + 5 = %d\n", reply)

	err = client.Call("Calculator.Multiply", args, &reply)
	if err != nil {
		log.Fatalf("Error calling Calculator.Multiply: %v", err)
	}
	fmt.Printf("10 * 5 = %d\n", reply)
}
