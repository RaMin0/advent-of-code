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
	var ops0 []int
	for _, opStr := range opsStr {
		op, err := strconv.Atoi(opStr)
		if err != nil {
			log.Fatalf("Failed to parse op %q: %v", opStr, err)
		}
		ops0 = append(ops0, op)
	}

	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			var ops []int
			for _, op := range ops0 {
				ops = append(ops, op)
			}

			ops[1], ops[2] = n, v

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

			if ops[0] == 19690720 {
				log.Printf("Result: %v", 100*n+v)
				break
			}
		}
	}
}
