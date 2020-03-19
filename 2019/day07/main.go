package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type computer struct {
	ops, ins, outs []int
}

func (c *computer) compute() error {
	for i := 0; i < len(c.ops); {
		modeop := c.ops[i]
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
				n = c.ops[n]
			case 1:
			default:
				log.Fatalf("Invalid mode: %d", mode[i])
			}
			return n
		}

		switch op {
		case 1:
			op1, op2, resIdx := opVal(c.ops[i+1], 1), opVal(c.ops[i+2], 2), c.ops[i+3]
			c.ops[resIdx] = op1 + op2
			i += 4
		case 2:
			op1, op2, resIdx := opVal(c.ops[i+1], 1), opVal(c.ops[i+2], 2), c.ops[i+3]
			c.ops[resIdx] = op1 * op2
			i += 4
		case 3:
			resIdx := c.ops[i+1]
			c.ops[resIdx] = c.ins[0]
			c.ins = c.ins[1:]
			i += 2
		case 4:
			op := opVal(c.ops[i+1], 1)
			c.outs = append([]int{op}, c.outs...)
			i += 2
		case 5:
			op1, op2 := opVal(c.ops[i+1], 1), opVal(c.ops[i+2], 2)
			if op1 != 0 {
				i = op2
			} else {
				i += 3
			}
		case 6:
			op1, op2 := opVal(c.ops[i+1], 1), opVal(c.ops[i+2], 2)
			if op1 == 0 {
				i = op2
			} else {
				i += 3
			}
		case 7:
			op1, op2, resIdx := opVal(c.ops[i+1], 1), opVal(c.ops[i+2], 2), c.ops[i+3]
			if op1 < op2 {
				c.ops[resIdx] = 1
			} else {
				c.ops[resIdx] = 0
			}
			i += 4
		case 8:
			op1, op2, resIdx := opVal(c.ops[i+1], 1), opVal(c.ops[i+2], 2), c.ops[i+3]
			if op1 == op2 {
				c.ops[resIdx] = 1
			} else {
				c.ops[resIdx] = 0
			}
			i += 4
		default:
			return fmt.Errorf("invalid op: %d", op)
		}
	}
	return nil
}

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

	var max int
	for _, cseq := range perm([]int{0, 1, 2, 3, 4}) {
		lastIn := 0
		for _, in := range cseq {
			c := computer{ops: make([]int, len(ops)), ins: []int{in, lastIn}}
			copy(c.ops, ops)
			if err := c.compute(); err != nil {
				log.Fatalf("Failed to compute: %v", err)
			}
			lastIn = c.outs[0]
		}
		if lastIn > max {
			max = lastIn
		}
	}
	log.Printf("Result: %v", max)
}

// Ref: https://rosettacode.org/wiki/Permutations#non-recursive.2C_lexicographical_order
func perm(a []int) (rtn [][]int) {
	for i, j, n, c := 0, 0, len(a)-1, 1; c < 120; c++ {
		i = n - 1
		j = n
		for a[i] > a[i+1] {
			i--
		}
		for a[j] < a[i] {
			j--
		}
		a[i], a[j] = a[j], a[i]
		j = n
		i++
		for i < j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
		rtna := make([]int, len(a))
		copy(rtna, a)
		rtn = append(rtn, rtna)
	}
	return
}
