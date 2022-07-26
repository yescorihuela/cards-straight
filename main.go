// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
	"math"
)

type ConfigCards struct {
	straightCards int
	maxCards      int
}

func main() {
	// input := []int{14, 2, 3, 4, 5, 6}
	// input := []int{14, 2, 3, 4, 5, 6}
	input := []int{7, 7, 12, 11, 3, 4, 14}
	c := NewConfigCards(5, 7)
	fmt.Println(c.isStraight(input))
}

func NewConfigCards(straightCards, maxCards int) ConfigCards {
	if straightCards == 0 || maxCards == 0 {

		return ConfigCards{}
	}
	if straightCards > maxCards {
		return ConfigCards{}
	}
	return ConfigCards{
		straightCards: straightCards,
		maxCards:      maxCards,
	}
}

func (c *ConfigCards) isStraight(input []int) bool {

	if len(input) < c.straightCards {
		log.Println("Input length is less than staircase cards config...")
		return false
	}
	if len(input) > c.maxCards {
		log.Println("Input length is greater than maximum cards config...")
		return false
	}
	cardsTranslator := map[int]int{
		14: 1,
		2:  2,
		3:  3,
		4:  4,
		5:  5,
		6:  6,
		7:  7,
		8:  8,
		9:  9,
		10: 10,
		11: 11,
		12: 12,
		13: 13,
	}
	stackCheckedCards := make([]int, 0)
	counter := 0
	for i, card := range input {
		stackCheckedCards = append(stackCheckedCards, cardsTranslator[card])
		if i > 0 && len(stackCheckedCards) > 0 {
			fmt.Println(i, stackCheckedCards)
			diff := cardsTranslator[card] - stackCheckedCards[i-1]
			if math.Abs(float64(diff)) == 1.0 {
				counter += 1
			} else {
				counter = 0
			}
			if counter == c.straightCards {
				return true
			}
		}
	}
	return false
}
