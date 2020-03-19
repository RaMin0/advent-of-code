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

	// ops[1], ops[2] = 12, 2

	for i := 0; i < len(ops); {
		modeop := ops[i]
		if modeop == 99 {
			break
		}

		modeInt, op := modeop/100, modeop%100

		opVal := func(n, i int) int {
			mode := make([]int, 3, 3)
			opModeInt := modeInt
			for m := 0; opModeInt > 0; m++ {
				mode[m] = opModeInt % 10
				opModeInt /= 10
			}

			switch mode[i-1] {
			case 0:
				n = ops[n]
			case 1:
			default:
				log.Fatalf("Invalid mode: %d", mode[i])
			}
			return n
		}

		switch op {
		case 1:
			op1, op2, resIdx := opVal(ops[i+1], 1), opVal(ops[i+2], 2), ops[i+3]
			ops[resIdx] = op1 + op2
			i += 4
		case 2:
			op1, op2, resIdx := opVal(ops[i+1], 1), opVal(ops[i+2], 2), ops[i+3]
			ops[resIdx] = op1 * op2
			i += 4
		case 3:
			resIdx := ops[i+1]
			ops[resIdx] = 1
			i += 2
		case 4:
			op := opVal(ops[i+1], 1)
			if op != 0 {
				log.Printf("Result: %v", op)
			}
			i += 2
		default:
			log.Fatalf("Invalid op: %d", op)
		}
	}
}
