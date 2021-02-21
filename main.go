// files in the same package can freely call functions from each other
package main

func main() {
	cards := newDeckFromFile("all_cards.txt")
	cards.print()
}
