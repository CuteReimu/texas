package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	turn       int32
	players    []IPlayer
	whoseTurn  int32
	whoIsFirst int32
	maxIn      int32
	totalIn    int32
}

func main() {
	const totalPlayerCount = 8
	game := &Game{}
	for i := 0; i < totalPlayerCount; i++ {
		if i == 0 {
			game.players = append(game.players, &CmdPlayer{})
		} else {
			game.players = append(game.players, &RobotPlayer{})
		}
	}
	for i := 0; i < totalPlayerCount; i++ {
		game.players[i].InitGame(totalPlayerCount, 10000)
	}
	random := rand.NewSource(time.Now().UnixNano() / 1000000)
	for {
		game.turn++
		fmt.Printf("第%d场\n", game.turn)

		game.OneGame(random)

		if len(game.players[0].GetAllCard()) < 7 {
			game.NotifyResult1()
		} else {
			game.NotifyResult()
		}
		fmt.Println("====================")
		if game.turn >= 100000000 {
			break
		}
		game.whoIsFirst++
		game.whoIsFirst %= int32(len(game.players))
	}
}

func (game *Game) OneGame(random rand.Source) {
	fmt.Println("每人摸了2张牌")
	deck := NewDeck(random)
	for _, player := range game.players {
		deck.Draw(1)
		cards := deck.Draw(2)
		player.Init(cards)
	}
	game.maxIn = 0
	game.totalIn = 0
	game.whoseTurn = game.whoIsFirst
	game.NextPlayer(false)
	game.NotifyIn(1, false)
	game.NextPlayer(false)
	game.NotifyIn(2, false)
	game.NextPlayer(true)
	game.maxIn = 2
	game.totalIn = 3
	if game.LoopIn() {
		return
	}

	deck.Draw(1)
	cards := deck.Draw(3)
	game.NotifyOpen(cards)
	game.whoseTurn = game.whoIsFirst
	game.NextPlayer(true)
	if game.LoopIn() {
		return
	}

	deck.Draw(1)
	cards = deck.Draw(1)
	game.NotifyOpen(cards)
	game.whoseTurn = game.whoIsFirst
	game.NextPlayer(true)
	if game.LoopIn() {
		return
	}

	deck.Draw(1)
	cards = deck.Draw(1)
	game.NotifyOpen(cards)
	game.whoseTurn = game.whoIsFirst
	game.NextPlayer(true)
	game.LoopIn()
}

func (game *Game) LoopIn() (isWin bool) {
	for {
		isWin = true
		for id, player := range game.players {
			if int32(id) != game.whoseTurn && !player.IsPassed() {
				isWin = false
				break
			}
		}
		beforeInCount := game.players[game.whoseTurn].GetInCount()
		if isWin || game.players[game.whoseTurn].IsAlreadyIn() && beforeInCount == game.maxIn {
			return
		}
		// 规则：加注不得超过场上已有总注数
		canInMaxCount := game.maxIn + game.totalIn
		count := game.players[game.whoseTurn].AskForIn(canInMaxCount)
		game.NotifyIn(count, true)
		if count > game.maxIn {
			game.maxIn = count
		}
		if count > 0 {
			game.totalIn += count - beforeInCount
		}
		game.NextPlayer(true)
	}
}

func (game *Game) NextPlayer(notify bool) {
	totalCount := int32(len(game.players))
	for {
		game.whoseTurn++
		game.whoseTurn %= totalCount
		if !game.players[game.whoseTurn].IsPassed() {
			break
		}
	}
	if notify {
		for id, player := range game.players {
			player.NotifyTurn((game.whoseTurn + totalCount - int32(id)) % totalCount)
		}
	}
}

func (game *Game) NotifyIn(count int32, isManual bool) {
	if count == 0 {
		fmt.Printf("%d号玩家弃牌\n", game.whoseTurn)
	} else if count == game.maxIn {
		if count == game.players[game.whoseTurn].GetInCount() {
			fmt.Printf("%d号玩家不加注\n", game.whoseTurn)
		} else {
			fmt.Printf("%d号玩家跟注\n", game.whoseTurn)
		}
	} else {
		fmt.Printf("%d号玩家加注至%d\n", game.whoseTurn, count)
	}
	totalCount := int32(len(game.players))
	for id, player := range game.players {
		player.NotifyIn((game.whoseTurn+totalCount-int32(id))%totalCount, count, isManual)
	}
}

func (game *Game) NotifyOpen(cards []*Card) {
	var logString []interface{}
	logString = append(logString, "翻出了")
	for _, card := range cards {
		logString = append(logString, card)
	}
	fmt.Println(logString...)
	for _, player := range game.players {
		player.NotifyOpen(cards)
	}
}

func (game *Game) NotifyResult1() {
	winnerId := int32(-1)
	for id, player := range game.players {
		if !player.IsPassed() {
			if winnerId > 0 {
				panic(winnerId)
			}
			winnerId = int32(id)
		}
	}
	if winnerId < 0 {
		panic(winnerId)
	}
	for id := range game.players {
		if winnerId == int32(id) {
			fmt.Println("总计筹码有：", game.totalIn, "赢家是：", winnerId)
			game.NotifyAddCoin(int32(id), game.totalIn)
		}
	}
}

func (game *Game) NotifyResult() {
	var results []*CheckResult
	for i, player := range game.players {
		if player.IsPassed() {
			results = append(results, nil)
		} else {
			var logString []interface{}
			result := Check(player.GetAllCard())
			results = append(results, result)
			logString = append(logString, "玩家", i, "的手牌是")
			for _, card := range player.GetHandCard() {
				logString = append(logString, card)
			}
			logString = append(logString, "，牌型是", result)
			fmt.Println(logString...)
		}
	}
	index := MaxResult(results...)
	buff, _ := json.Marshal(index)
	fmt.Println("总计筹码有：", game.totalIn, "赢家是：", string(buff))
	j := int32(1)
	for id := range game.players {
		isWin := false
		for _, i := range index {
			if i == int32(id) {
				isWin = true
				break
			}
		}
		//stat.Add(player.GetHandCard(), isWin, len(index))
		if isWin {
			game.NotifyAddCoin(int32(id), (game.totalIn-j)/int32(len(index))+1)
			j++
		}
	}
}

func (game *Game) NotifyAddCoin(playerId, count int32) {
	fmt.Println("玩家", playerId, "获得", count, "筹码（+", count-game.players[playerId].GetInCount(), "）")
	totalCount := int32(len(game.players))
	for id, player := range game.players {
		player.NotifyAddCoin((playerId+totalCount-int32(id))%totalCount, count)
	}
}
