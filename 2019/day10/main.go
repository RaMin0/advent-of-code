package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type p struct{ x, y int }
type v struct{ x, y float64 }

func main() {
	var asts []p
	var w, h int
	sc := bufio.NewScanner(os.Stdin)
	for y := 0; sc.Scan(); y++ {
		h++
		for x, c := range strings.TrimSpace(sc.Text()) {
			w++
			if c == '.' {
				continue
			}
			asts = append(asts, p{x, y})
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	w /= h

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

	sees := make([]map[v]bool, len(asts), len(asts))
	for i, src := range asts {
		for _, dst := range asts {
			dx, dy := dst.x-src.x, dst.y-src.y

			if dx == 0 && dy == 0 {
				continue
			}

			dgcd := float64(gcd(dx, dy))
			dstv := v{float64(dx) / dgcd, float64(dy) / dgcd}

			if sees[i] == nil {
				sees[i] = map[v]bool{}
			}
			sees[i][dstv] = true
		}
	}

	sort.Slice(sees, func(i, j int) bool { return len(sees[i]) > len(sees[j]) })

	log.Printf("Result: %v", len(sees[0]))
}
