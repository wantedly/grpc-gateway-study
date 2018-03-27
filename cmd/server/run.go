package main

import (
	"os"

	"github.com/wantedly/grpc-gateway-study/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := app.Run()
	if err != nil {
		return 1
	}
	return 0
}
