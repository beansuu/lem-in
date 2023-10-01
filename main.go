package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AntNest struct {
	ants    int
	rooms   map[string][2]int //x, y
	tunnels map[string][]string
	start   string
	end     string
}

func (an AntNest) addRoom(name string, x, y int) {
	an.rooms[name] = [2]int{x, y}
}

func (an AntNest) addTunnel(room1, room2 string) {
	an.tunnels[room1] = append(an.tunnels[room1], room2)
	an.tunnels[room2] = append(an.tunnels[room2], room1)
}

func parseAntNest(filename string) AntNest {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	antNest := AntNest{0, make(map[string][2]int), make(map[string][]string), "", ""}
	scanner := bufio.NewScanner(file)
	// read number of ants
	if !scanner.Scan() {
		panic("No ants")
	}
	antNest.ants = parseNum(scanner.Text())
	// read rooms and tunnels
	for scanner.Scan() {
		switch scanner.Text() {
		case "##start":
			if !scanner.Scan() {
				panic("No start room")
			}
			name, x, y := parseRoom(scanner.Text())
			antNest.addRoom(name, x, y)
			antNest.start = name
		case "##end":
			if !scanner.Scan() {
				panic("No end room")
			}
			name, x, y := parseRoom(scanner.Text())
			antNest.addRoom(name, x, y)
			antNest.end = name
		default:
			if strings.HasPrefix(scanner.Text(), "#") {
				continue
			}
			parts := strings.Split(scanner.Text(), " ")
			if len(parts) == 3 {
				name, x, y := parseRoom(scanner.Text())
				antNest.addRoom(name, x, y)
			} else {
				room1, room2 := parseTunnel(scanner.Text())
				antNest.addTunnel(room1, room2)
			}
		}
	}
	if antNest.ants == 0 || len(antNest.rooms) == 0 || len(antNest.tunnels) == 0 || antNest.start == "" || antNest.end == "" {
		panic("Invalid ant nest")
	}
	return antNest
}

func parseRoom(s string) (string, int, int) {
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		panic("Invalid room")
	}
	return parts[0], parseNum(parts[1]), parseNum(parts[2])
}

func parseTunnel(s string) (string, string) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		panic("Invalid tunnel")
	}
	return parts[0], parts[1]
}

func parseNum(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	if len(os.Args) < 2 {
		println("Usage: go run main.go <filename>")
		return
	}
	antNest := parseAntNest(os.Args[1])
	fmt.Printf("%+v\n", antNest)
}
