package main

import (
	"os"
)

func main() {
	name := os.Args[1]
	if name != "" {
		RecursiveTouch(name)
	}
}
