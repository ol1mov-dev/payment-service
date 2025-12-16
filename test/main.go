package main

import (
	"fmt"
	"slices"
)

type Car interface {
	Name() string
}

func (car *Car) Name() string {
	fmt.Println()
}

func main() {
	s := []string{
		"one", "two", "lolololol",
	}

	_ = []string{
		"three", "four", "five", "three",
	}

	b := slices.Repeat(s, 123)

	fmt.Println(b)
}
