package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	service := NewUsersService()

	rpc.Register(service)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatal("listen error:", err)
	}

	http.Serve(listener, nil)
}
