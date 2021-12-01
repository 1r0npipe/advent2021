package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var counter int
	arr := make([]int, 0, 2003)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Can't get interger from string")
		}
		arr = append(arr, current)
	}
	var sum, current int
	for i := 0; i < len(arr)-3; i++ {
		sum = arr[i] + arr[i+1] + arr[i+2]
		if sum > current {
			counter++
		}
		current = sum
	}
	fmt.Printf("This is the result %d\n", counter)
}
