package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			eachCard := value + " of " + suit
			cards = append(cards, eachCard)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) printDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deckToString() string {
	res := strings.Join([]string(d), ",")
	return res
}

func stringToDeck(rawDeck []byte) deck {
	return deck(strings.Split(string(rawDeck), ","))
}

func (d deck) saveDeckToFile(fileName string) error {
	err := ioutil.WriteFile(fileName, []byte(d.deckToString()), 0666)
	return err
}

func readDeckFromFile(fileName string) deck {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return stringToDeck(content)
}

func (d deck) shuffleCards() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for ind := range d {
		randomNumber := r.Intn(len(d) - 1)
		d[ind], d[randomNumber] = d[randomNumber], d[ind]
	}
	return d
}
