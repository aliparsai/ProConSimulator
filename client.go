package main

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/kaptinlin/jsonschema"
)

func runClient(port string) error {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		return fmt.Errorf("Error connecting: %w", err)
	}
	defer conn.Close()

	data, err := ioutil.ReadAll(conn)
	if err != nil {
		return fmt.Errorf("Error reading: %w", err)
	}

	schemaData, err := ioutil.ReadFile("schema.json")
	if err != nil {
		return fmt.Errorf("Error reading schema: %w", err)
	}

	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile(schemaData)
	if err != nil {
		return fmt.Errorf("Error compiling schema: %w", err)
	}

	result := schema.Validate(data)
	if result.IsValid() {
		fmt.Println("Received data is valid.")
		var person Person
		if err := schema.Unmarshal(&person, data); err == nil {
			fmt.Printf("Received Person: %+v\n", person)
		}
		return nil
	}

	fmt.Println("Received data is invalid.")
	for field, err := range result.Errors {
		fmt.Printf("- %s: %s\n", field, err.Message)
	}
	return fmt.Errorf("validation failed")
}
