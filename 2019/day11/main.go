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
	pc, rbase      int
}

func (c *computer) ensure(i int) int {
	if i >= len(c.ops) {
		ops := make([]int, i+1, i+1)
		copy(ops, c.ops)
		c.ops = ops
	}
	return i
}

func (c *computer) read(i int) int {
	return c.ops[c.ensure(i)]
}

func (c *computer) write(i, v int) {
	c.ops[c.ensure(i)] = v
}

func (c *computer) compute() error {
	for {
		modeop := c.ops[c.pc]
		if modeop == 99 {
			return errHalt
		}

		modeInt, op := modeop/100, modeop%100
		mode := make([]int, 3, 3)
		for m := 0; modeInt > 0; m++ {
			mode[m] = modeInt % 10
			modeInt /= 10
		}

		read := func(n, i int) int {
			switch mode[i-1] {
			case 0:
				n = c.read(n)
			case 1:
			case 2:
				n = c.read(c.rbase + n)
			default:
				log.Fatalf("invalid mode: %d", mode[i-1])
			}
			return n
		}

		res := func(n, i int) int {
			switch mode[i-1] {
			case 0:
			case 1:
			case 2:
				n += c.rbase
			default:
				log.Fatalf("invalid mode: %d", mode[i-1])
			}
			return n
		}

		switch op {
		case 1:
			op1, op2, resIdx := read(c.ops[c.pc+1], 1), read(c.ops[c.pc+2], 2), res(c.ops[c.pc+3], 3)
			c.write(resIdx, op1+op2)
			c.pc += 4
		case 2:
			op1, op2, resIdx := read(c.ops[c.pc+1], 1), read(c.ops[c.pc+2], 2), res(c.ops[c.pc+3], 3)
			c.write(resIdx, op1*op2)
			c.pc += 4
		case 3:
			if len(c.ins) == 0 {
				return nil
			}
			resIdx := res(c.ops[c.pc+1], 1)
			c.write(resIdx, c.ins[0])
			c.ins = c.ins[1:]
			c.pc += 2
		case 4:
			op := read(c.ops[c.pc+1], 1)
			c.outs = append([]int{op}, c.outs...)
			c.pc += 2
		case 5:
			op1, op2 := read(c.ops[c.pc+1], 1), read(c.ops[c.pc+2], 2)
			if op1 != 0 {
				c.pc = op2
			} else {
				c.pc += 3
			}
		case 6:
			op1, op2 := read(c.ops[c.pc+1], 1), read(c.ops[c.pc+2], 2)
			if op1 == 0 {
				c.pc = op2
			} else {
				c.pc += 3
			}
		case 7:
			op1, op2, resIdx := read(c.ops[c.pc+1], 1), read(c.ops[c.pc+2], 2), res(c.ops[c.pc+3], 3)
			if op1 < op2 {
				c.write(resIdx, 1)
			} else {
				c.write(resIdx, 0)
			}
			c.pc += 4
		case 8:
			op1, op2, resIdx := read(c.ops[c.pc+1], 1), read(c.ops[c.pc+2], 2), res(c.ops[c.pc+3], 3)
			if op1 == op2 {
				c.write(resIdx, 1)
			} else {
				c.write(resIdx, 0)
			}
			c.pc += 4
		case 9:
			op := read(c.ops[c.pc+1], 1)
			c.rbase += op
			c.pc += 2
		default:
			return fmt.Errorf("invalid op: %d", op)
		}
	}
}

type p struct{ x, y int }

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

	pos, dir := p{0, 0}, p{0, 1}
	grid := map[p]int{}
	turn := func(d int) {
		if d == 1 {
			switch dir {
			case p{0, 1}:
				dir = p{1, 0}
			case p{1, 0}:
				dir = p{0, -1}
			case p{0, -1}:
				dir = p{-1, 0}
			case p{-1, 0}:
				dir = p{0, 1}
			}
			return
		}
		switch dir {
		case p{0, 1}:
			dir = p{-1, 0}
		case p{-1, 0}:
			dir = p{0, -1}
		case p{0, -1}:
			dir = p{1, 0}
		case p{1, 0}:
			dir = p{0, 1}
		}
	}
	move := func() {
		pos = p{pos.x + dir.x, pos.y + dir.y}
	}

	c := computer{ops: ops}

	for {
		c.ins = append(c.ins, grid[pos])
		if err := c.compute(); err != nil {
			if err == errHalt {
				break
			}
			log.Fatalf("Failed to compute: %v", err)
		}
		grid[pos] = c.outs[1]
		turn(c.outs[0])
		move()
	}
	log.Printf("Result: %v", len(grid))
}
