package main

import (
	"fmt"

	"github.com/alexchungnc/project1/backend/function/cmd/server"
)

func main() {
	fmt.Println("hello from maim.")
	server.Serve()
}
