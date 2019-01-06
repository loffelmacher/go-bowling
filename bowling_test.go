package main

import "testing"

func RollMany(g *BowlingGame, nn, pins int) {
	for ii := 0; ii < nn; ii++ {
		g.Roll(pins)
	}
}

func TestGutterGame(t *testing.T) {
	g := BowlingGame{}
	RollMany(&g, 20, 0)
	if g.Score() != 0 {
		t.Fail()
	}
}

func TestAllOnes(t *testing.T) {
	g := BowlingGame{}
	RollMany(&g, 20, 1)
	if g.Score() != 20 {
		t.Fail()
	}
}

func TestOneSpare(t *testing.T) {
	g := BowlingGame{}
	g.Roll(5)
	g.Roll(5)
	g.Roll(3)
	RollMany(&g, 17, 0)
	if g.Score() != 16 {
		t.Log("Expected 16 but score is " + g.String())
		t.Fail()
	}
}
