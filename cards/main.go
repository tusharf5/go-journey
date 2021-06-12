package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 3)

	// hand.saveToFile("name.txt")
	newhand := readFromFile("name.txt")

	fmt.Println(newhand.toString())

}
