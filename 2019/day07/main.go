package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var errHalt = fmt.Errorf("halt")

type computer struct {
	ops, ins, outs []int
	pc             int
}

func (c *computer) compute() error {
	for {
		modeop := c.ops[c.pc]
		if modeop == 99 {
			return errHalt
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
			op1, op2, resIdx := opVal(c.ops[c.pc+1], 1), opVal(c.ops[c.pc+2], 2), c.ops[c.pc+3]
			c.ops[resIdx] = op1 + op2
			c.pc += 4
		case 2:
			op1, op2, resIdx := opVal(c.ops[c.pc+1], 1), opVal(c.ops[c.pc+2], 2), c.ops[c.pc+3]
			c.ops[resIdx] = op1 * op2
			c.pc += 4
		case 3:
			resIdx := c.ops[c.pc+1]
			c.ops[resIdx] = c.ins[0]
			c.ins = c.ins[1:]
			c.pc += 2
		case 4:
			op := opVal(c.ops[c.pc+1], 1)
			c.outs = append([]int{op}, c.outs...)
			c.pc += 2
			return nil
		case 5:
			op1, op2 := opVal(c.ops[c.pc+1], 1), opVal(c.ops[c.pc+2], 2)
			if op1 != 0 {
				c.pc = op2
			} else {
				c.pc += 3
			}
		case 6:
			op1, op2 := opVal(c.ops[c.pc+1], 1), opVal(c.ops[c.pc+2], 2)
			if op1 == 0 {
				c.pc = op2
			} else {
				c.pc += 3
			}
		case 7:
			op1, op2, resIdx := opVal(c.ops[c.pc+1], 1), opVal(c.ops[c.pc+2], 2), c.ops[c.pc+3]
			if op1 < op2 {
				c.ops[resIdx] = 1
			} else {
				c.ops[resIdx] = 0
			}
			c.pc += 4
		case 8:
			op1, op2, resIdx := opVal(c.ops[c.pc+1], 1), opVal(c.ops[c.pc+2], 2), c.ops[c.pc+3]
			if op1 == op2 {
				c.ops[resIdx] = 1
			} else {
				c.ops[resIdx] = 0
			}
			c.pc += 4
		default:
			return fmt.Errorf("invalid op: %d", op)
		}
	}
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
	for _, cseq := range perm([]int{5, 6, 7, 8, 9}) {
		computers := make([]*computer, len(cseq))
		for i, in := range cseq {
			c := computer{ops: make([]int, len(ops)), ins: []int{in}}
			copy(c.ops, ops)
			computers[i] = &c
		}

		halted := map[*computer]bool{}
		lastIn := 0
		for {
			for _, c := range computers {
				if halted[c] {
					continue
				}
				c.ins = append(c.ins, lastIn)
				if err := c.compute(); err != nil {
					if err == errHalt {
						halted[c] = true
					} else {
						log.Fatalf("Failed to compute: %v", err)
					}
				}
				lastIn = c.outs[0]
			}
			if len(halted) == len(cseq) {
				break
			}
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
