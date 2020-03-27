package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

var src = p{19, 11}

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

	angle := func(dx, dy float64) (int, float64) {
		switch {
		case dx == 0 && dy < 0:
			return 0, 0
		case dx > 0 && dy < 0:
			return 1, dx / math.Abs(dy)
		case dx > 0 && dy == 0:
			return 2, 0
		case dx > 0 && dy > 0:
			return 3, dy / dx
		case dx == 0 && dy > 0:
			return 4, 0
		case dx < 0 && dy > 0:
			return 5, math.Abs(dx) / dy
		case dx < 0 && dy == 0:
			return 6, 0
		case dx < 0 && dy < 0:
			return 7, dy / dx
		}
		return -1, -1
	}

	anglesMap := map[[2]float64][]p{}
	for _, dst := range asts {
		dx, dy := dst.x-src.x, dst.y-src.y

		if dx == 0 && dy == 0 {
			continue
		}

		dgcd := float64(gcd(dx, dy))
		dstv := v{float64(dx) / dgcd, float64(dy) / dgcd}
		s, o := angle(dstv.x, dstv.y)
		a := [2]float64{float64(s), o}
		anglesMap[a] = append(anglesMap[a], dst)
	}

	dist := func(m, n p) float64 { return math.Sqrt(math.Pow(float64(n.x-m.x), 2) + math.Pow(float64(n.y-m.y), 2)) }

	var angles [][2]float64
	for a, asts := range anglesMap {
		angles = append(angles, a)
		sort.Slice(anglesMap[a], func(i, j int) bool { return dist(src, asts[i]) < dist(src, asts[j]) })
	}
	sort.Slice(angles, func(i, j int) bool {
		return angles[i][0] < angles[j][0] || (angles[i][0] == angles[j][0] && angles[i][1] < angles[j][1])
	})

	for v := 0; ; {
		for _, a := range angles {
			if len(anglesMap) == 0 {
				return
			}

			if anglesMap[a] == nil {
				continue
			}

			var ast p
			ast, anglesMap[a] = anglesMap[a][0], anglesMap[a][1:]
			v++
			if v == 200 {
				log.Printf("Result: %v", 100*ast.x+ast.y)
				return
			}
			if len(anglesMap[a]) == 0 {
				delete(anglesMap, a)
			}
		}
	}
}
