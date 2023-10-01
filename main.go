package main

import (
	"ant/ant"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: go run main.go <filename>")
		return
	}
	antNest := ant.ParseNest(os.Args[1])
	fmt.Printf("%+v\n", antNest)
}
