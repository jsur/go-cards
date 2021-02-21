// files in the same package can freely call functions from each other
package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
