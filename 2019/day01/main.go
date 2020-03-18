package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	var totalFuel int

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; sc.Scan(); i++ {
		line := sc.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to parse line %d %s : %v", i, line, err)
		}
		fuel := calcFuel(mass)
		totalFuel += fuel
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Total fuel: %d", totalFuel)
}

func calcFuel(mass int) int {
	fuel := mass/3 - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + calcFuel(fuel)
}
