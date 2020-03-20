package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	w, h := 25, 6
	area := w * h
	var layers []string
	for c := 0; c < len(bytes.TrimSpace(b)); {
		var l string
		for i := 1; i <= area; i, c = i+1, c+1 {
			l += string(b[c])
		}
		layers = append(layers, l)
	}

	image := []byte(layers[0])
	for i := 0; i < len(image); i++ {
		for l := 1; l < len(layers); l++ {
			if image[i] == '2' {
				image[i] = layers[l][i]
			}
		}
	}

	for i, p := range image {
		if i > 0 && i%w == 0 {
			fmt.Println()
		}
		if p == '0' {
			p = ' '
		}
		fmt.Print(string(p))
	}
	fmt.Println()
}
