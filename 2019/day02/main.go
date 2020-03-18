package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	opsStr := strings.Split(strings.TrimSpace(string(b)), ",")
	var ops []int
	for _, opStr := range opsStr {
		op, err := strconv.Atoi(opStr)
		if err != nil {
			log.Fatalf("Failed to parse op %q: %v", opStr, err)
		}
		ops = append(ops, op)
	}

	ops[1], ops[2] = 12, 2

	for i := 0; i < len(ops); i += 4 {
		op := ops[i]
		if op == 99 {
			break
		}

		op1, op2, resIdx := ops[ops[i+1]], ops[ops[i+2]], ops[i+3]
		switch op {
		case 1:
			ops[resIdx] = op1 + op2
		case 2:
			ops[resIdx] = op1 * op2
		default:
			log.Fatalf("Invalid op: %d", op)
		}
	}

	log.Printf("Result: %v", ops[0])
}
