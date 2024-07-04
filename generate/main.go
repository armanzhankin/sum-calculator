package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	rand.NewSource(time.Now().UnixNano())
	numObjects := 1000000
	data := make([]Data, numObjects)

	for i := range data {
		data[i] = Data{
			A: rand.Intn(21) - 10, // [-10, 10]
			B: rand.Intn(21) - 10, // [-10, 10]
		}
	}

	file, err := os.Create("data.json")
	if err != nil {
		log.Fatalf("Couldn't create generated JSON: %s", err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		log.Fatalf("Couldn't encode generated JSON: %s", err.Error())
	}
}
