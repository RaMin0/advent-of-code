package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"math"
	"os"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	area := 25 * 6
	minZeros, minZerosIdx := math.MaxInt64, -1
	var layers []string
	for c := 0; c < len(bytes.TrimSpace(b)); {
		var l string
		var zeros int
		for i := 1; i <= area; i++ {
			if b[c] == '0' {
				zeros++
			}
			l += string(b[c])
			c++
		}
		if zeros < minZeros {
			minZeros = zeros
			minZerosIdx = len(layers)
		}
		layers = append(layers, l)
	}

	var ones, twos int
	for _, c := range layers[minZerosIdx] {
		if c == '1' {
			ones++
		}
		if c == '2' {
			twos++
		}
	}
	log.Printf("Result: %v", ones*twos)
}
