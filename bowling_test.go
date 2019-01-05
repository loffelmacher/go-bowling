package main

import "testing"

func TestGutterGame(t *testing.T) {
	g := BowlingGame{}
	for ii := 0; ii < 20; ii++ {
		g.Roll(0)
	}
	if g.Score() != 0 {
		t.Fail()
	}
}

func TestAllOnes(t *testing.T) {
	g := BowlingGame{}
	for ii := 0; ii < 20; ii++ {
		g.Roll(1)
	}
	if g.Score() != 20 {
		t.Fail()
	}
}
