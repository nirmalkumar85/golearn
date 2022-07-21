package main

import "fmt"

//added comment

func main() {

	cards := newDeck()
	cards.print()

	fmt.Println("-------------------")
	//fmt.Println(cards.toString())
	//fmt.Println([]byte(cards.toString()))

	//cards.saveToFile("my_cards")

	//newCards := newDeckFromFile("new_cards")
	//newCards.print()

	cards.shuffleDeck()
	//cards.print()
}
