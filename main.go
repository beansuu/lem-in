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

	shortestPath := antNest.FindShortestPath()
	fmt.Println(len(shortestPath))
	for i, room := range shortestPath {
		fmt.Printf("L%d-%s ", i+1, room)
	}
	fmt.Println() // Print a new line at the end

	// fmt.Printf("%+v\n", antNest)
}
