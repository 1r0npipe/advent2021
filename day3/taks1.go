package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)


func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic(err)
	}
	arr := make([]string, 0, 1000)
	var gamma, epsilon string
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}
	var neg, pos int
	for i := 0; i < len(arr[0]); i++ {
		neg = 0
		pos = 0
		for _, bit := range arr {
			if bit[i] == '0' {
				neg += 1
			} else {
				pos += 1
			}
		}
		if pos > neg {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}
	output, err := strconv.ParseInt(gamma, 2, 64)  
	if err != nil {  
		log.Fatal("Could not make int from binary")
	}
	output1, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {  
		log.Fatal("Could not make int from binary")
	}  
	fmt.Printf("Result = %d\n", output * output1)
}

