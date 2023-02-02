package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name string
	var timeout string
	var namespace string

	flag.StringVar(&name, "name", "", "The name of the application")
	flag.StringVar(&timeout, "timeout", "", "The timeout in seconds to wait for the result")
	flag.StringVar(&namespace, "namespace", "", "The namespace of the application")

	flag.Parse()

	if name == "" {
		fmt.Println("Error: name flag must be set.")
		os.Exit(1)
	}

	if timeout == "" {
		fmt.Println("Error: timeout flag must be set.")
		os.Exit(1)
	}

	if namespace == "" {
		fmt.Println("Error: namespace flag must be set.")
		os.Exit(1)
	}

	fmt.Println("Name was set to: " + name)
	fmt.Println("Timeout was set to: " + timeout)
	fmt.Println("Namespace was set to: " + namespace)
}
