package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type step struct {
	dir string
	num int
}

func main() {
	var allSteps [][]step
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		var steps []step
		line := sc.Text()
		stepsStr := strings.Split(line, ",")

		for _, stepStr := range stepsStr {
			dir := stepStr[:1]
			num, err := strconv.Atoi(stepStr[1:])
			if err != nil {
				fmt.Printf("Failed to parse %q: %v\n", stepStr, err)
				return
			}
			steps = append(steps, step{dir: dir, num: num})
		}
		allSteps = append(allSteps, steps)
	}
	if err := sc.Err(); err != nil {
		fmt.Println(err)
		return
	}

	coords := map[[2]int]map[int]bool{}
	for i, steps := range allSteps {
		var x, y int
		for _, step := range steps {
			delta := move(&x, &y, step.dir, step.num)
			for _, coord := range delta {
				if _, ok := coords[coord]; !ok {
					coords[coord] = map[int]bool{}
				}
				coords[coord][i] = true
			}
		}
	}

	var coordDists []int
	for coord, n := range coords {
		if len(n) <= 1 {
			continue
		}
		coordDists = append(coordDists, dist(coord))
	}

	sort.Ints(coordDists)
	fmt.Printf("Result: %v\n", coordDists[0])
}

func move(x, y *int, dir string, num int) [][2]int {
	var coords [][2]int
	for ; num > 0; num-- {
		switch dir {
		case "R":
			*x++
		case "L":
			*x--
		case "U":
			*y++
		case "D":
			*y--
		}
		coords = append(coords, [2]int{*x, *y})
	}
	return coords
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func dist(coord [2]int) int {
	return abs(coord[0]) + abs(coord[1])
}
