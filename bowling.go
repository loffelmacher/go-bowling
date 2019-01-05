package main

import "fmt"

type BowlingGame struct {
	score int
}

func (g *BowlingGame) Score() int {
	return g.score
}

func (g *BowlingGame) Roll(pins int) {
	g.score += pins
}

func main() {
	fmt.Println("Hello world")
}
