package main

import (
	"fmt"
	"log"
	"os"
)

type CmdPlayer struct {
	totalCoin int32
	cards     []*Card
	players   []playerState
	maxIn     int32
}

func (r *CmdPlayer) InitGame(playerCount, coinCount int32) {
	r.players = make([]playerState, playerCount)
	r.totalCoin = coinCount
}

func (r *CmdPlayer) Init(cards []*Card) {
	if r.totalCoin > 0 {
		fmt.Println("你现在有", r.totalCoin, "筹码")
	} else {
		fmt.Println("很遗憾，你输光了")
		_, _ = fmt.Scanln()
		os.Exit(0)
	}
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

func (r *CmdPlayer) NotifyTurn(_ int32) {
}

func (r *CmdPlayer) AskForIn(maxCount int32) int32 {
	fmt.Println()
	var totalCount int32
	var logString []interface{}
	logString = append(logString, "当前你已下注：", r.GetInCount(), "，其他玩家的投注额为：")
	for id, player := range r.players {
		totalCount += player.inCount
		if id > 0 {
			if player.passed {
				logString = append(logString, "弃")
			} else {
				logString = append(logString, player.inCount)
			}
		}
	}
	fmt.Println(logString...)
	fmt.Println("池内的总筹码数为：", totalCount)
	logString = []interface{}{}
	logString = append(logString, "你的手牌为")
	cards := r.GetAllCard()
	logString = append(logString, cards[0], cards[1])
	if len(cards) > 2 {
		logString = append(logString, "桌上亮出的牌为")
	}
	for i := 2; i < len(cards); i++ {
		logString = append(logString, cards[i])
	}
	fmt.Println(logString...)
	if !r.IsAlreadyAdd() {
		fmt.Println("请输入你想要加注至多少，0表示弃牌，", r.maxIn, "表示跟注/不加注，大于", r.maxIn, "的数表示加注至多少，最高加注至", maxCount)
		for {
			var a int32
			n, err := fmt.Scanln(&a)
			if err != nil || n != 1 || !(a == 0 || a >= r.maxIn && a <= maxCount) {
				if err != nil {
					log.Println(err)
				}
				fmt.Println("你的输入有误，请重新输入")
				continue
			}
			return a
		}
	} else {
		fmt.Println("请输入你想要加注至多少，0表示弃牌，", r.maxIn, "表示跟注/不加注，因为你本轮已经加注过，所以不能加注")
		for {
			var a int32
			n, err := fmt.Scanln(&a)
			if err != nil || n != 1 || !(a == 0 || a == r.maxIn) {
				if err != nil {
					log.Println(err)
				}
				fmt.Println("你的输入有误，请重新输入")
				continue
			}
			return a
		}
	}
}

func (r *CmdPlayer) NotifyIn(playerId, count int32, isManual bool) {
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

func (r *CmdPlayer) NotifyOpen(cards []*Card) {
	for i := range r.players {
		r.players[i].alreadyIn = false
		r.players[i].alreadyAdd = false
	}
	for _, card := range cards {
		r.cards = append(r.cards, &Card{card.Color, card.Num})
	}
}

func (r *CmdPlayer) NotifyAddCoin(playerId, count int32) {
	if playerId == 0 {
		r.totalCoin += count
	}
}

func (r *CmdPlayer) GetHandCard() []*Card {
	return r.cards[:2:2]
}

func (r *CmdPlayer) GetAllCard() []*Card {
	return r.cards
}

func (r *CmdPlayer) GetInCount() int32 {
	return r.players[0].inCount
}

func (r *CmdPlayer) GetTotalCoin() int32 {
	return r.totalCoin
}

func (r *CmdPlayer) IsPassed() bool {
	return r.players[0].passed
}

func (r *CmdPlayer) IsAlreadyAdd() bool {
	return r.players[0].alreadyAdd
}

func (r *CmdPlayer) IsAlreadyIn() bool {
	return r.players[0].alreadyIn
}
