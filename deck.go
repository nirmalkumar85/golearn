package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "DIamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {

		for _, cardValue := range cardValues {

			cards = append(cards, cardValue+" of "+cardSuit)

		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {

	//take a deck to a string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {

		//Option 1 - log the error and call to newDeck()
		//Option 2 - log the error and quit

		fmt.Println("Error:", err)
		os.Exit(1)

	}

	deckString := string(bs[:])
	newDeck := strings.Split(deckString, ",")

	return newDeck
}

func (d deck) shuffleDeck() {

	d.print()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })

	fmt.Println("--------------")
	d.print()

}
