package main

import (
	"encoding/json"
	"sort"
)

type CheckResult struct {
	CardType int32
	Nums     []int
}

func (result *CheckResult) String() string {
	buf, _ := json.Marshal(result.Nums)
	switch result.CardType {
	case 9:
		return "同花顺" + string(buf)
	case 8:
		return "四条" + string(buf)
	case 7:
		return "葫芦" + string(buf)
	case 6:
		return "同花" + string(buf)
	case 5:
		return "顺子" + string(buf)
	case 4:
		return "三条" + string(buf)
	case 3:
		return "两对" + string(buf)
	case 2:
		return "一对" + string(buf)
	case 1:
		return "无" + string(buf)
	}
	panic(result.CardType)
}

func Check(cards []*Card) *CheckResult {
	var result *CheckResult
	result = checkStraightFlush(cards)
	if result != nil {
		return result
	}
	result = checkFourOfAKind(cards)
	if result != nil {
		return result
	}
	result = checkFullHouse(cards)
	if result != nil {
		return result
	}
	result = checkFlush(cards)
	if result != nil {
		return result
	}
	result = checkStraight(cards)
	if result != nil {
		return result
	}
	result = checkThreeOfAKind(cards)
	if result != nil {
		return result
	}
	return checkPair(cards)
}

func checkStraightFlush(cards []*Card) *CheckResult {
	data := make([][]int, 5)
	for _, card := range cards {
		data[card.Color] = append(data[card.Color], int(card.Num))
	}
	for _, d := range data {
		if len(d) >= 5 {
			sort.Ints(d)
			if len(d) == 7 {
				if isStraight(d[2:]) {
					return &CheckResult{9, d[2:]}
				} else if isStraight(d[1:6]) {
					return &CheckResult{9, d[1:6]}
				} else if isStraight(d[:5]) {
					return &CheckResult{9, d[:5]}
				}
			} else if len(d) == 6 {
				if isStraight(d[1:]) {
					return &CheckResult{9, d[1:]}
				} else if isStraight(d[:5]) {
					return &CheckResult{9, d[:5]}
				}
			} else {
				if isStraight(d) {
					return &CheckResult{9, d}
				}
			}
			if d[0] == 2 && d[1] == 3 && d[2] == 4 && d[3] == 5 && d[len(d)-1] == 14 {
				return &CheckResult{9, []int{1, 2, 3, 4, 5}}
			}
		}
	}
	return nil
}

func checkFourOfAKind(cards []*Card) *CheckResult {
	data := make([]int, 0, len(cards))
	for _, card := range cards {
		data = append(data, int(card.Num))
	}
	sort.Ints(data)
	count := 1
	for i := 1; i < len(data); i++ {
		if data[i] == data[i-1] {
			count++
		} else {
			count = 1
		}
		if count >= 4 {
			if i == len(data)-1 {
				return &CheckResult{8, []int{data[i-4], data[i], data[i], data[i], data[i]}}
			} else {
				return &CheckResult{8, []int{data[len(data)-1], data[i], data[i], data[i], data[i]}}
			}
		}
	}
	return nil
}

func checkFullHouse(cards []*Card) *CheckResult {
	count := make([]int, 15)
	for _, card := range cards {
		count[card.Num]++
	}
	var card3, card2 int
	for i := 2; i <= 14; i++ {
		switch count[i] {
		case 3:
			if card3 != 0 {
				card2 = card3
			}
			card3 = i
		case 2:
			card2 = i
		}
	}
	if card3 != 0 && card2 != 0 {
		return &CheckResult{7, []int{card2, card2, card3, card3, card3}}
	}
	return nil
}

func checkFlush(cards []*Card) *CheckResult {
	data := make([][]int, 5)
	for _, card := range cards {
		data[card.Color] = append(data[card.Color], int(card.Num))
	}
	for _, d := range data {
		if len(d) >= 5 {
			sort.Ints(d)
			return &CheckResult{6, d[len(d)-5:]}
		}
	}
	return nil
}

func checkStraight(cards []*Card) *CheckResult {
	count := make([]int, 15)
	for _, card := range cards {
		count[card.Num]++
	}
	for i := 10; i >= 2; i-- {
		if count[i] > 0 && count[i+1] > 0 && count[i+2] > 0 && count[i+3] > 0 && count[i+4] > 0 {
			return &CheckResult{5, []int{i, i + 1, i + 2, i + 3, i + 4}}
		}
	}
	if count[14] > 0 && count[2] > 0 && count[3] > 0 && count[4] > 0 && count[5] > 0 {
		return &CheckResult{5, []int{1, 2, 3, 4, 5}}
	}
	return nil
}

func checkThreeOfAKind(cards []*Card) *CheckResult {
	data := make([]int, 0, len(cards))
	for _, card := range cards {
		data = append(data, int(card.Num))
	}
	sort.Ints(data)
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}
	count := 1
	for i := 1; i < len(data); i++ {
		if data[i] == data[i-1] {
			count++
		} else {
			count = 1
		}
		if count >= 3 {
			data2 := make([]int, 0, len(cards)-3)
			for j := 0; j < len(data); j++ {
				if data[j] != data[i] {
					data2 = append(data2, data[j])
				}
			}
			return &CheckResult{4, []int{data2[1], data2[0], data[i], data[i], data[i]}}
		}
	}
	return nil
}

func checkPair(cards []*Card) *CheckResult {
	count := make([]int, 15)
	for _, card := range cards {
		count[card.Num]++
	}
	var card1, card2 []int
	for i := 14; i >= 2; i-- {
		if count[i] == 2 {
			card2 = append(card2, i)
		} else if count[i] == 1 {
			card1 = append(card1, i)
		}
	}
	if len(card2) >= 2 && len(card1) >= 1 {
		c := card1[0]
		if len(card2) == 3 && card2[2] > c {
			c = card2[2]
		}
		return &CheckResult{3, []int{c, card2[1], card2[1], card2[0], card2[0]}}
	}
	if len(card2) == 1 && len(card1) >= 3 {
		return &CheckResult{2, []int{card1[2], card1[1], card1[0], card2[0], card2[0]}}
	}
	if len(card2) == 0 && len(card1) >= 5 {
		return &CheckResult{1, []int{card1[4], card1[3], card1[2], card1[1], card1[0]}}
	}
	panic(cards)
}

func MaxResult(results ...*CheckResult) []int32 {
	var maxIndex []int32
	totalPlayerCount := int32(len(results))
	for i := int32(0); i < totalPlayerCount; i++ {
		if results[i] == nil {
			continue
		}
		if len(maxIndex) == 0 {
			maxIndex = []int32{i}
		} else {
			switch results[maxIndex[0]].Compare(results[i]) {
			case -1:
				maxIndex = []int32{i}
			case 0:
				maxIndex = append(maxIndex, i)
			}
		}
	}
	return maxIndex
}

func (result *CheckResult) Compare(result2 *CheckResult) int {
	if result.CardType > result2.CardType {
		return 1
	} else if result.CardType < result2.CardType {
		return -1
	}
	for i := 4; i >= 0; i-- {
		if result.Nums[i] > result2.Nums[i] {
			return 1
		} else if result.Nums[i] < result2.Nums[i] {
			return -1
		}
	}
	return 0
}

func isStraight(d []int) bool {
	if len(d) != 5 {
		panic(d)
	}
	return d[0]+1 == d[1] && d[1]+1 == d[2] && d[2]+1 == d[3] && d[3]+1 == d[4]
}
