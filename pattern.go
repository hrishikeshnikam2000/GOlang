package main

import "fmt"

func main() {
	var num = 10
	lowerLeftCorner := num*(num-1)/2 + 1
	lastInColumn := lowerLeftCorner
	lastInRow := 1
	for i, row := 1, 1; row <= num; i++ {
		w := len(fmt.Sprint(lastInColumn))
		if i < lastInRow {
			fmt.Printf("%*d ", w, i)
			lastInColumn++
		} else {
			fmt.Printf("%*d\n", w, i)
			row++
			lastInRow += row
			lastInColumn = lowerLeftCorner
		}
	}
}