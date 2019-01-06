package main

import "fmt"

type BowlingGame struct {
	rolls       [21]int
	currentRoll int
}

func (g *BowlingGame) Score() int {
	var score = 0
	var frameIndex = 0
	for frame := 0; frame < 10; frame++ {
		if g.rolls[frameIndex] == 10 {
			score += 10 + g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
			frameIndex++
		} else if g.isSpare(frameIndex) {
			score += 10 + g.rolls[frameIndex+2]
			frameIndex += 2
		} else {
			score += g.rolls[frameIndex] + g.rolls[frameIndex+1]
			frameIndex += 2
		}

	}
	return score
}

func (g *BowlingGame) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

func (g *BowlingGame) String() string {
	return fmt.Sprintf("%d", g.Score())
}

func (g *BowlingGame) isSpare(frameIndex int) bool {
	return g.rolls[frameIndex]+g.rolls[frameIndex+1] == 10
}

func main() {
	fmt.Println("Hello world")
}
