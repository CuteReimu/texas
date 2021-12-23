package main

import (
	"math"
	"time"
)

type IPlayer interface {
	InitGame(playerCount, coinCount int32)
	Init(cards []*Card)
	NotifyTurn(playerId int32)
	AskForIn(maxCount int32) int32
	NotifyIn(playerId, count int32, isManual bool)
	NotifyOpen(cards []*Card)
	NotifyAddCoin(playerId, count int32)
	GetHandCard() []*Card
	GetAllCard() []*Card
	GetInCount() int32
	GetTotalCoin() int32
	IsPassed() bool
	IsAlreadyAdd() bool
	IsAlreadyIn() bool
}

type playerState struct {
	inCount    int32
	passed     bool
	alreadyAdd bool
	alreadyIn  bool
}

type RobotPlayer struct {
	totalCoin int32
	cards     []*Card
	players   []playerState
	maxIn     int32
}

func (r *RobotPlayer) InitGame(playerCount, coinCount int32) {
	r.players = make([]playerState, playerCount)
	r.totalCoin = coinCount
}

func (r *RobotPlayer) Init(cards []*Card) {
	r.cards = make([]*Card, 2, 7)
	r.cards[0] = &Card{cards[0].Color, cards[0].Num}
	r.cards[1] = &Card{cards[1].Color, cards[1].Num}
	for i := range r.players {
		r.players[i].inCount = 0
		r.players[i].passed = false
		r.players[i].alreadyAdd = false
		r.players[i].alreadyIn = false
	}
	r.maxIn = 2
}

func (r *RobotPlayer) NotifyTurn(_ int32) {
}

func (r *RobotPlayer) AskForIn(maxCount int32) int32 {
	var count int32
	switch len(r.cards) {
	case 2:
		count = r.calFirstIn(maxCount)
	case 5:
		count = r.calSecondIn(maxCount)
	case 6:
		count = r.calThirdIn(maxCount)
	case 7:
		count = r.calLastIn(maxCount)
	default:
		panic(len(r.cards))
	}
	if count > maxCount {
		panic(count)
	}
	time.Sleep(time.Second / 2)
	return count
}

func (r *RobotPlayer) NotifyIn(playerId, count int32, isManual bool) {
	if isManual {
		r.players[playerId].alreadyIn = true
	}
	if count == 0 {
		r.players[playerId].passed = true
	} else {
		addCount := count - r.players[playerId].inCount
		if addCount > 0 {
			if isManual {
				r.players[playerId].alreadyAdd = true
			}
			r.players[playerId].inCount = count
			r.maxIn = count
			if playerId == 0 {
				r.totalCoin -= addCount
			}
		}
	}
}

func (r *RobotPlayer) NotifyOpen(cards []*Card) {
	for i := range r.players {
		r.players[i].alreadyIn = false
		r.players[i].alreadyAdd = false
	}
	for _, card := range cards {
		r.cards = append(r.cards, &Card{card.Color, card.Num})
	}
}

func (r *RobotPlayer) NotifyAddCoin(playerId, count int32) {
	if playerId == 0 {
		r.totalCoin += count
	}
}

func (r *RobotPlayer) GetHandCard() []*Card {
	return r.cards[:2:2]
}

func (r *RobotPlayer) GetAllCard() []*Card {
	return r.cards
}

func (r *RobotPlayer) GetInCount() int32 {
	return r.players[0].inCount
}

func (r *RobotPlayer) GetTotalCoin() int32 {
	return r.totalCoin
}

func (r *RobotPlayer) IsPassed() bool {
	return r.players[0].passed
}

func (r *RobotPlayer) IsAlreadyAdd() bool {
	return r.players[0].alreadyAdd
}

func (r *RobotPlayer) IsAlreadyIn() bool {
	return r.players[0].alreadyIn
}

func (r *RobotPlayer) calFirstIn(maxCount int32) int32 {
	// 首注直接读概率表
	winRate := stat.Get(r.GetHandCard())
	playerCount := int32(len(r.players))
	var inPlayerCount int32
	for i := range r.players {
		if r.players[i].alreadyIn || r.players[i].passed {
			inPlayerCount++
		}
	}
	if inPlayerCount > playerCount-3 {
		inPlayerCount = playerCount - 3
	}
	needWinRate := 0.2 - 0.1*float64(inPlayerCount)/float64(playerCount-3)
	if winRate < needWinRate && r.GetInCount() < r.maxIn {
		if r.GetInCount() == r.maxIn {
			return r.maxIn
		} else {
			return 0
		}
	}
	if r.IsAlreadyAdd() {
		return r.maxIn
	}
	if r.maxIn > 2+int32(len(r.players))/2 {
		if winRate > 0.2 && winRate >= needWinRate+0.05 {
			return maxCount
		}
	} else {
		if winRate >= needWinRate+0.05 {
			return maxCount
		}
	}
	return r.maxIn
}

func (r *RobotPlayer) calSecondIn(maxCount int32) int32 {
	// 算概率和赔率
	hashList := CardListToHash(r.GetAllCard())
	var cards []*Card
	for _, card := range r.GetAllCard() {
		cards = append(cards, &Card{card.Color, card.Num})
	}
	cards = append(cards, &Card{0, 0}, &Card{0, 0})
	winCount := 0
	totalCount := 0
	for i := CardHashMin; i <= CardHashMax; i++ {
		if inList(i, hashList) {
			continue
		}
		cards[5] = GetCard(i)
		for j := CardHashMin; j <= CardHashMax; j++ {
			if i == j || inList(j, hashList) {
				continue
			}
			cards[6] = GetCard(j)
			totalCount++
			if Check(cards).CardType >= 5 {
				winCount++
			}
		}
	}
	winRate := float64(winCount) / float64(totalCount)
	return r.calRate(winRate, maxCount)
}

func (r *RobotPlayer) calThirdIn(maxCount int32) int32 {
	// 算概率和赔率
	hashList := CardListToHash(r.GetAllCard())
	var cards []*Card
	for _, card := range r.GetAllCard() {
		cards = append(cards, &Card{card.Color, card.Num})
	}
	cards = append(cards, &Card{0, 0})
	winCount := 0
	totalCount := 0
	for i := CardHashMin; i <= CardHashMax; i++ {
		if inList(i, hashList) {
			continue
		}
		cards[6] = GetCard(i)
		totalCount++
		if Check(cards).CardType >= 5 {
			winCount++
		}
	}
	winRate := float64(winCount) / float64(totalCount)
	return r.calRate(winRate, maxCount)
}

func (r *RobotPlayer) calLastIn(maxCount int32) int32 {
	// 算概率和赔率
	cards0 := r.GetAllCard()
	checkResult := Check(r.GetAllCard())
	hashList := CardListToHash(r.GetAllCard())
	var cards []*Card
	for i := 2; i < 7; i++ {
		cards = append(cards, &Card{cards0[i].Color, cards0[i].Num})
	}
	cards = append(cards, &Card{0, 0}, &Card{0, 0})
	winCount := 0
	totalCount := 0
	for i := CardHashMin; i <= CardHashMax; i++ {
		if inList(i, hashList) {
			continue
		}
		cards[5] = GetCard(i)
		for j := CardHashMin; j <= CardHashMax; j++ {
			if i == j || inList(j, hashList) {
				continue
			}
			cards[6] = GetCard(j)
			checkResult2 := Check(cards)
			if checkResult2.CardType >= 5 {
				totalCount += 2
				compareResult := checkResult.Compare(checkResult2)
				if compareResult >= 0 {
					winCount++
				}
				if compareResult == 0 {
					winCount++
				}
			}
		}
	}
	winRate := float64(winCount) / float64(totalCount)
	var leftPlayerCount int32
	for _, player := range r.players {
		if !player.passed {
			leftPlayerCount++
		}
	}
	return r.calRate(math.Pow(winRate, float64(leftPlayerCount-1)), maxCount)
}

func inList(a int32, l []int32) bool {
	for _, i := range l {
		if i == a {
			return true
		}
	}
	return false
}

func (r *RobotPlayer) calRate(winRate float64, maxCount int32) int32 {
	var leftPlayerCount, totalCoin, passCoin int32
	for _, player := range r.players {
		totalCoin += player.inCount
		if !player.passed {
			leftPlayerCount++
		} else {
			passCoin += player.inCount
		}
	}
	rate := float64(r.maxIn-r.GetInCount()) / float64(totalCoin)
	if winRate < rate {
		return 0
	}
	if r.IsAlreadyAdd() {
		return r.maxIn
	}
	if 1/winRate-float64(leftPlayerCount) < 0 {
		return maxCount
	}
	x := (float64(r.GetInCount())/winRate + float64(passCoin)) / (1/winRate - float64(leftPlayerCount))
	if x > float64(maxCount) {
		return maxCount
	}
	if x > float64(r.maxIn)+float64(len(r.players)) {
		return int32(x) - int32(len(r.players))
	}
	return r.maxIn
}
