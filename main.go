package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Numbers struct {
	A int `json:"a"`
	B int `json:"b"`
}

func fileRead(filename string) ([]Numbers, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %s", err.Error())
	}

	var numbers []Numbers
	if err := json.Unmarshal(file, &numbers); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}

	return numbers, nil
}

func sumNumbers(numbers []Numbers, goroutinesNumber int) int {
	sliceLen := len(numbers)
	chunkSize := (sliceLen + goroutinesNumber - 1) / goroutinesNumber

	resultChan := make(chan int, goroutinesNumber)
	var wg sync.WaitGroup

	for i := 0; i < goroutinesNumber; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > sliceLen {
			end = sliceLen
		}
		wg.Add(1)
		go worker(numbers[start:end], resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	return totalSum
}

func worker(numbers []Numbers, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, num := range numbers {
		sum += num.A + num.B
	}
	resultChan <- sum
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Please run as: 'go run main.go <path-to-json> <goroutines-number>'")
	}

	jsonFile := os.Args[1]
	goroutinesNumber, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Failed to parse number of goroutines, please enter integer")
	}

	numbers, err := fileRead(jsonFile)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	totalSum := sumNumbers(numbers, goroutinesNumber)
	fmt.Printf("Sum of Numbers: %d\n", totalSum)
}
