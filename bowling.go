package main

import "fmt"

type BowlingGame struct {
	Score int
}

func (g *BowlingGame) score() int {
	return g.Score
}

func (g *BowlingGame) roll(pins int) {
	g.Score += pins
}

func main() {
	fmt.Println("Hello world")
}
