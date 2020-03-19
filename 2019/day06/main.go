package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type node struct {
	up   *node
	down []*node
}

func (n *node) addChild(c *node) {
	c.up = n
	n.down = append(n.down, c)
}

func main() {
	nodes := map[string]*node{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		parts := strings.Split(strings.TrimSpace(sc.Text()), ")")
		parentStr, childStr := parts[0], parts[1]

		parent, child := &node{}, &node{}
		if _, ok := nodes[parentStr]; !ok {
			nodes[parentStr] = parent
		}
		if _, ok := nodes[childStr]; !ok {
			nodes[childStr] = child
		}
		nodes[parentStr].addChild(nodes[childStr])
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int
	for _, n := range nodes {
		cn := *n
		for cn.up != nil {
			sum++
			cn = *cn.up
		}
	}
	log.Printf("Result: %v", sum)
}
