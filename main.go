package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	idx int
	x   int
	y   int
}

type Link struct {
	room1 int
	room2 int
}

func main() {
	if len(os.Args) < 2 {
		println("Usage: go run main.go <filename>")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rooms []Room
	var links []Link
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch scanner.Text() {
		case "##start":
			fallthrough
		case "##end":
			fallthrough
		default:
			if !strings.HasPrefix(scanner.Text(), "#") {
				parts := strings.Split(scanner.Text(), " ")
				if len(parts) > 2 {
					idx := parseNum(parts[0])
					x := parseNum(parts[1])
					y := parseNum(parts[2])
					room := Room{idx, x, y}
					rooms = append(rooms, room)
				} else {
					parts = strings.Split(scanner.Text(), "-")
					link := Link{parseNum(parts[0]), parseNum(parts[1])}
					links = append(links, link)
				}
			}
		}
	}
	fmt.Println("Rooms", rooms)
	fmt.Println("Links", links)
}

func parseNum(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
