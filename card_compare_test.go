package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	cards := []*Card{
		{2, 5},
		{1, 14},
		{1, 10},
		{1, 8},
		{1, 2},
		{3, 9},
		{2, 9},
	}
	t.Log(cards)
	t.Log(Check(cards))
}
