package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	line, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	inputsStr := strings.Split(strings.TrimSpace(string(line)), "-")
	lower, err := strconv.Atoi(inputsStr[0])
	if err != nil {
		log.Fatalf("Failed to parse input %q: %v", inputsStr[0], err)
	}
	upper, err := strconv.Atoi(inputsStr[1])
	if err != nil {
		log.Fatalf("Failed to parse input %q: %v", inputsStr[1], err)
	}

	checks := []func(string) bool{
		func(n string) bool { return len(n) == 6 },
		func(n string) bool {
			for i := 1; i < len(n); i++ {
				c0, c := n[i-1:i], n[i:i+1]
				if c == c0 {
					return true
				}
				c0 = c
			}
			return false
		},
		func(n string) bool {
			for i := 1; i < len(n); i++ {
				c0, c := n[i-1], n[i]
				if c < c0 {
					return false
				}
			}
			return true
		},
	}

	var valid int
loop:
	for n := lower; n <= upper; n++ {
		nStr := strconv.Itoa(n)
		for _, check := range checks {
			if !check(nStr) {
				continue loop
			}
		}
		valid++
	}
	log.Printf("Result: %v\n", valid)
}
