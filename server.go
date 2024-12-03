package main

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Calculator struct{}

func (c *Calculator) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func (c *Calculator) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {
	calculator := new(Calculator)
	err := rpc.Register(calculator)
	if err != nil {
		log.Fatalf("Error registering Calculator: %v", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	log.Println("RPC server is running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
