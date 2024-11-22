package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

func main() {
	// setup flags
	var bootstrapServers = flag.StringArray("bootstrap_server", []string{}, "bootstrap servers to connect to")
	flag.Parse()

	// run
	fmt.Println("hello", bootstrapServers)
}
