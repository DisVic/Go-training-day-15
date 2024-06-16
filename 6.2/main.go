package main

import (
	"fmt"
)

func main() {
	var c rune
	_, _ = fmt.Scanf("%c", &c)
	if c <= 'Z' && c >= 'A' {
		c += 32
	} else {
		c -= 32
	}
	fmt.Printf("%c", c)
}
