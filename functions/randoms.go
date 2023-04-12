package functions

import "math/rand"

func NewRandomValue(upper int) int {
	randomValue := rand.Intn(upper)
	return randomValue
}

func RandomIndex(vector []int) int {
	return rand.Intn(len(vector))
}
