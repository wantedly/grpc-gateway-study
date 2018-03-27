package main

import (
	"fmt"
	"os"

	"github.com/wantedly/grpc-gateway-study/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := app.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "an error occured: %v", err)
		return 1
	}
	return 0
}
