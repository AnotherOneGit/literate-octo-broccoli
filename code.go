package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var homeBet int
var visitorBet int
var homeScore int
var visitorScore int
var homeWin = false
var visitorWin = false
var draw = false

func main() {
	talk()
	score()
	result()

	switch {
	case homeWin:
		fmt.Println("Хозяева победили!")
	case visitorWin:
		fmt.Println("Гости победили!")
	case draw:
		fmt.Println("Ничья!")
	}
}

func talk() {
	// спрашиваем у пользователя и выводим прогноз счёта
	fmt.Print("Введите количество голов хозяев: ")
	fmt.Fscan(os.Stdin, &homeBet)

	fmt.Print("Введите количество голов гостей: ")
	fmt.Fscan(os.Stdin, &visitorBet)
	fmt.Println(fmt.Sprintf("Вы поставили на счёт %d : %d", homeBet, visitorBet))
}

func score() {
	//формируем и выводим рандомный счёт с вариантами от нуля до трёх
	rand.Seed(time.Now().UnixNano())
	score := [2]int{rand.Intn(4), rand.Intn(4)}

	homeScore = score[0]
	visitorScore = score[1]
	fmt.Println(fmt.Sprintf("Результат матча - %d : %d", homeScore, visitorScore))
}

func result() {
	// определения результата матча
	if homeScore > visitorScore {
		homeWin = true
	} else if homeScore < visitorScore {
		visitorWin = true
	}
	draw = true
}
