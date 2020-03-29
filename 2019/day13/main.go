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

	ops[0] = 2

	c := computer{ops: ops}

	var score, paddlex, ballx int
	grid := map[p]int{}
	for {
		err := c.compute()
		if err != nil && err != errHalt {
			log.Fatalf("Failed to compute: %v", err)
		}

		for i := 0; i < len(c.outs); i += 3 {
			t, y, x := c.outs[i], c.outs[i+1], c.outs[i+2]
			if x == -1 {
				score = t
				continue
			}

			grid[p{x, y}] = t

			switch t {
			case 3:
				paddlex = x
			case 4:
				ballx = x
			}
		}
		c.outs = nil

		// draw(grid)
		// fmt.Printf("Score: %v\n\n", score)

		if err == errHalt {
			break
		}

		// time.Sleep(50 * time.Millisecond)

		var in int
		switch {
		case paddlex < ballx:
			in = 1
		case paddlex > ballx:
			in = -1
		}
		c.ins = append(c.ins, in)
	}

	log.Printf("Result: %v", score)
}

func draw(grid map[p]int) {
	maxx, maxy := -1, -1
	for c := range grid {
		if c.x > maxx {
			maxx = c.x
		}
		if c.y > maxy {
			maxy = c.y
		}
	}

	var sb strings.Builder
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			switch t := grid[p{x, y}]; t {
			case 0:
				fmt.Fprint(&sb, " ")
			case 1:
				fmt.Fprint(&sb, "#")
			case 2:
				fmt.Fprint(&sb, "x")
			case 3:
				fmt.Fprint(&sb, "\"")
			case 4:
				fmt.Fprint(&sb, "o")
			default:
				log.Fatalf("invalid tile: %d", t)
			}
		}
		fmt.Fprintln(&sb)
	}
	fmt.Println(sb.String())
}
