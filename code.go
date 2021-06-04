package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var homeBet, visitorBet, homeScore, visitorScore int
var homeWin, visitorWin, draw bool

type match struct {
	home, visitor int
	result        string
}

func main() {
	bet()
	score()

	switch {
	case homeWin:
		fmt.Println("Хозяева победили!")
	case visitorWin:
		fmt.Println("Гости победили!")
	case draw:
		fmt.Println("Ничья!")
	}
}

func bet() {
	// спрашиваем у пользователя и выводим прогноз счёта
	fmt.Print("Введите количество голов хозяев: ")
	fmt.Fscan(os.Stdin, &homeBet)

	fmt.Print("Введите количество голов гостей: ")
	fmt.Fscan(os.Stdin, &visitorBet)

	bet := match{
		home:    homeBet,
		visitor: visitorBet,
		result:  result(homeBet, visitorBet),
	}

	fmt.Println(fmt.Sprintf("Вы поставили на счёт %d : %d", bet.home, bet.visitor))
}

func score() {
	//формируем и выводим рандомный счёт с вариантами от нуля до трёх
	rand.Seed(time.Now().UnixNano())
	score := [2]int{rand.Intn(4), rand.Intn(4)}

	homeScore = score[0]
	visitorScore = score[1]

	game := match{
		home:    homeScore,
		visitor: visitorScore,
		result:  result(homeScore, visitorScore),
	}

	fmt.Println(fmt.Sprintf("Результат матча - %d : %d", game.home, game.visitor))

}

func result(home int, visitor int) string {
	// определения результата матча
	if home > visitor {
		return "Хозяева победили!"
	} else if home < visitor {
		return "Гости победили!"
	}
	return "Ничья"
}
