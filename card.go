package main

import (
	"math/rand"
	"strconv"
)

type Card struct {
	Color int32
	Num   int32
}

const (
	CardHashMin int32 = 1*13 + 2
	CardHashMax int32 = 4*13 + 14
)

func (card *Card) Hash() int32 {
	return card.Color*13 + card.Num
}

func CardListToHash(cards []*Card) []int32 {
	result := make([]int32, len(cards))
	for i, card := range cards {
		result[i] = card.Hash()
	}
	return result
}

func GetCard(hash int32) *Card {
	card := &Card{hash / 13, hash % 13}
	if card.Num < 2 {
		card.Num += 13
		card.Color--
	}
	return card
}

func (card *Card) String() string {
	var color, num string
	switch card.Color {
	case 1:
		color = "♣"
	case 2:
		color = "♦"
	case 3:
		color = "♥"
	case 4:
		color = "♠"
	}
	switch card.Num {
	case 11:
		num = "J"
	case 12:
		num = "Q"
	case 13:
		num = "K"
	case 14:
		num = "A"
	default:
		num = strconv.Itoa(int(card.Num))
	}
	return color + num
}

type Deck struct {
	cards []*Card
	pos   int
}

func NewDeck(random rand.Source) *Deck {
	d := new(Deck)
	for i := int32(1); i < 4; i++ {
		for j := int32(2); j <= 14; j++ {
			d.cards = append(d.cards, &Card{i, j})
		}
	}
	d.pos = len(d.cards)
	l := d.pos
	for i := 0; i < l; i++ {
		j := int(random.Int63())%(l-i) + i
		if i != j {
			d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
		}
	}
	return d
}

func (d *Deck) Draw(n int) []*Card {
	result := make([]*Card, 0, n)
	for i := 0; i < n; i++ {
		d.pos--
		result = append(result, d.cards[d.pos])
	}
	return result
}
