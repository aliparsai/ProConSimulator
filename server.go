package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func runServer(port string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listen.Close()
	fmt.Println("Server listening on :" + port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	person := Person{
		Name: "Jules",
		Age:  30,
	}

	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(person); err != nil {
		fmt.Println("Error encoding person:", err.Error())
	}
}
