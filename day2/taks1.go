package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	var hor, depth, depth1, aim int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tempSlice := strings.Split(line, " ")
		pos, err := strconv.Atoi(tempSlice[1])
		if err != nil {
			log.Fatal("Can't get interger from string")
		}
		switch {
		case tempSlice[0] == "forward":
			hor = hor + pos
			depth1 = depth1 + aim*pos
		case tempSlice[0] == "down":
			depth = depth + pos
			aim = aim + pos
		case tempSlice[0] == "up":
			depth = depth - pos
			aim = aim - pos
		}
	}
	fmt.Printf("Part1 result = %d\n", depth*hor)
	fmt.Printf("Part2 result = %d\n", depth1*hor)
}
