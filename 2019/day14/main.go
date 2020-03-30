package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Amount struct {
	Chemical string
	Quantity int
}

func (a Amount) String() string { return fmt.Sprintf("{%v %v}", a.Quantity, a.Chemical) }

func ParseAmount(raw string) Amount {
	rawArr := strings.Split(raw, " ")
	quantity, chemical := rawArr[0], rawArr[1]
	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		log.Fatal("unable to parse quantity %q: %v", quantity, err)
	}
	return Amount{chemical, quantityInt}
}

type Rule struct {
	Inputs []Amount
	Output Amount
}

func ParseRule(raw string) Rule {
	rawArr := strings.Split(raw, " => ")
	inputsStr, outputStr := rawArr[0], rawArr[1]
	var inputs []Amount
	for _, input := range strings.Split(inputsStr, ", ") {
		inputs = append(inputs, ParseAmount(input))
	}
	output := ParseAmount(outputStr)
	return Rule{inputs, output}
}

func LeastOre(rules []Rule, fuelRequired int) int {
	rulesByChemical := map[string]Rule{}
	for _, r := range rules {
		rulesByChemical[r.Output.Chemical] = r
	}

	requirements := map[string]int{"FUEL": fuelRequired}
	var oreNeeded int

	done := func() bool {
		for _, quantity := range requirements {
			if quantity > 0 {
				return false
			}
		}
		return true
	}

	for !done() {
		var key string
		for chemical, quantity := range requirements {
			if quantity <= 0 {
				continue
			}
			key = chemical
			break
		}
		quantityNeeded := requirements[key]

		rule := rulesByChemical[key]
		numTimes := int(math.Ceil(float64(quantityNeeded) / float64(rule.Output.Quantity)))
		requirements[key] -= numTimes * rule.Output.Quantity

		for _, a := range rule.Inputs {
			if a.Chemical == "ORE" {
				oreNeeded += numTimes * a.Quantity
				continue
			}
			requirements[a.Chemical] += numTimes * a.Quantity
		}
	}

	return oreNeeded
}

func MaxFuel(rules []Rule, maxOre int) int {
	return sort.Search(maxOre, func(i int) bool {
		return LeastOre(rules, i) > maxOre
	}) - 1
}

func main() {
	var rules []Rule
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		rules = append(rules, ParseRule(strings.TrimSpace(sc.Text())))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Result: %v", MaxFuel(rules, 1_000_000_000_000))
}
