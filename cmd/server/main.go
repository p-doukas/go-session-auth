package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: medapp [serve|migrate]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "serve":
		serve()
	case "migrate":
		migrate()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Usage: medapp [serve|migrate]")
		os.Exit(1)
	}
}
