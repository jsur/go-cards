// files in the same package can freely call functions from each other
package main

func main() {
	cards := newDeck()
	cards.saveToFile("all_cards.txt")
}
