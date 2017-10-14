package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"

	//only use := when assigning a new variable
	card := newCard()
	fmt.Println(card)
}

func newCard() string {

	return "Five of Diamonds"
}
