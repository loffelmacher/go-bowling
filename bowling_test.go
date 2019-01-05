package main

import "testing"

func TestGutterGame(t *testing.T) {
	g := BowlingGame{}
	for ii := 0; ii < 20; ii++ {
		g.roll(0)
	}
	if g.score() != 0 {
		t.Fail()
	}
}

func TestAllOnes(t *testing.T) {
	g := BowlingGame{}
	for ii := 0; ii < 20; ii++ {
		g.roll(1)
	}
	if g.score() != 20 {
		t.Fail()
	}
}
