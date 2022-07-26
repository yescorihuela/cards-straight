package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStraight(t *testing.T) {

	straightMock := NewConfigCards(5, 7)

	t.Run("When the input is correct", func(t *testing.T) {
		slicesTest := [][]int{
			{5, 6, 7, 8, 9, 11, 12},
			{14, 2, 3, 4, 5, 9, 10},
			{9, 10, 11, 12, 13, 14},
		}

		// a := []int{5, 6, 7, 8, 9, 11, 12}
		// b := []int{14, 2, 3, 4, 5, 9, 10}
		// c := []int{9, 10, 11, 12, 13, 14}
		// slicesTest = append(slicesTest, a, b, c)

		for _, input := range slicesTest {
			assert.True(t, straightMock.isStraight(input))
		}

	})

	t.Run("When the input is wrong", func(t *testing.T) {
		slicesTest := make([][]int, 0)

		a := []int{5, 6, 14, 8, 9, 11, 12}
		b := []int{14, 11, 3, 4, 2, 9, 10}
		c := []int{5, 10, 11, 6, 13, 14, 4}
		slicesTest = append(slicesTest, a, b, c)

		for _, input := range slicesTest {
			assert.False(t, straightMock.isStraight(input))
		}
	})
}
