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
		if isStrike(g, frameIndex) {
			score += 10 + strikeBonus(g, frameIndex)
			frameIndex++
		} else if isSpare(g, frameIndex) {
			score += 10 + spareBonus(g, frameIndex)
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

func spareBonus(g *BowlingGame, frameIndex int) int {
	return g.rolls[frameIndex+2]
}

func strikeBonus(g *BowlingGame, frameIndex int) int {
	return g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
}

func isSpare(g *BowlingGame, frameIndex int) bool {
	return g.rolls[frameIndex]+g.rolls[frameIndex+1] == 10
}

func isStrike(g *BowlingGame, frameIndex int) bool {
	return g.rolls[frameIndex] == 10
}

func main() {
	fmt.Println("Hello world")
}
