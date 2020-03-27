package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type p struct{ x, y, z int }

func (m *p) add(n p) { m.x, m.y, m.z = m.x+n.x, m.y+n.y, m.z+n.z }
func (m *p) sum() int {
	return int(math.Abs(float64(m.x)) + math.Abs(float64(m.y)) + math.Abs(float64(m.z)))
}

type moon struct{ pos, vel p }

func main() {
	var moons []moon
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		var m moon
		fmt.Sscanf(sc.Text(), "<x=%d, y=%d, z=%d>", &m.pos.x, &m.pos.y, &m.pos.z)
		moons = append(moons, m)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	for t := 1; t <= 1000; t++ {
		for i := range moons {
			for _, other := range moons {
				if moons[i].pos.x < other.pos.x {
					moons[i].vel.x++
				}
				if moons[i].pos.x > other.pos.x {
					moons[i].vel.x--
				}
				if moons[i].pos.y < other.pos.y {
					moons[i].vel.y++
				}
				if moons[i].pos.y > other.pos.y {
					moons[i].vel.y--
				}
				if moons[i].pos.z < other.pos.z {
					moons[i].vel.z++
				}
				if moons[i].pos.z > other.pos.z {
					moons[i].vel.z--
				}
			}
		}

		for i, m := range moons {
			moons[i].pos.add(m.vel)
		}
	}

	var e int
	for _, m := range moons {
		e += m.pos.sum() * m.vel.sum()
	}
	log.Printf("Result: %v", e)
}
