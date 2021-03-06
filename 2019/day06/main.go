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

	you := path(nodes["YOU"])[1:]
	san := path(nodes["SAN"])[1:]

	for y := 0; y < len(you); y++ {
		for s := 0; s < len(san); s++ {
			if you[y] == san[s] {
				ypath := you[:y]
				spath := san[:s]
				log.Printf("Result: %v", len(ypath)+len(spath))
				return
			}
		}
	}
}

func path(n *node) []*node {
	ns := []*node{}
	cn := n
	for {
		ns = append(ns, cn)
		if cn.up == nil {
			break
		}
		cn = cn.up
	}
	return ns
}
