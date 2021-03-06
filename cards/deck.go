package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//declare a type which is slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Ace", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Ace"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//receiver function associated with type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() deck {

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range d {
		pos := r.Intn(len(d) - 1)
		d[i], d[pos] = d[pos], d[i]
	}

	return d
}
