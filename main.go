package main

import (
	"flag"
	"fmt"
)

func main() {
	appType := flag.String("type", "server", "Application type: server or client")
	port := flag.String("port", "8080", "Port to use")
	flag.Parse()

	if *appType == "server" {
		runServer(*port)
	} else if *appType == "client" {
		if err := runClient(*port); err != nil {
			fmt.Println("Client error:", err)
		}
	} else {
		fmt.Println("Invalid application type. Use 'server' or 'client'.")
	}
}
