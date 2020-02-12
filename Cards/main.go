package main

import "fmt"

func main() {
	// Variable Declarations
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"

	// Function Declaration
	// card := newCard()

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
