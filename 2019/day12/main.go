package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type p [3]int

func (m *p) add(n p) { m[0], m[1], m[2] = m[0]+n[0], m[1]+n[1], m[2]+n[2] }
func (m *p) sum() int {
	return int(math.Abs(float64(m[0])) + math.Abs(float64(m[1])) + math.Abs(float64(m[2])))
}

type moon struct{ pos, vel p }

func main() {
	var moons [4]moon
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; sc.Scan(); i++ {
		var m moon
		fmt.Sscanf(sc.Text(), "<x=%d, y=%d, z=%d>", &m.pos[0], &m.pos[1], &m.pos[2])
		moons[i] = m
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	gcd := func(a, b int) int {
		abs := func(n int) int { return int(math.Abs(float64(n))) }
		min := func(n, m int) int { return int(math.Min(float64(m), float64(n))) }
		max := func(n, m int) int { return int(math.Max(float64(m), float64(n))) }

		a, b = abs(a), abs(b)
		small, gcd := min(a, b), max(a, b)
		for i := small; i > 0; i-- {
			if a%i == 0 && b%i == 0 {
				return i
			}
		}
		return gcd
	}

	count := func(c int) int {
		var ms [4]moon
		for mi, m := range moons {
			ms[mi] = m
		}
		history := map[[4]moon]bool{ms: true}
		for t := 1; ; t++ {
			for i := range moons {
				for _, other := range moons {
					if moons[i].pos[c] < other.pos[c] {
						moons[i].vel[c]++
					}
					if moons[i].pos[c] > other.pos[c] {
						moons[i].vel[c]--
					}
				}
			}

			for i, m := range moons {
				moons[i].pos.add(m.vel)
			}

			if history[moons] {
				return t
			}
			history[moons] = true
		}
	}

	x, y, z := count(0), count(1), count(2)
	xy := (x * y) / gcd(x, y)
	xyz := (xy * z) / gcd(xy, z)
	log.Printf("Result: %v", xyz)
}
