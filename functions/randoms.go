package functions

import "math/rand"

func NewRandomValue(upper int) int {
	randomValue := rand.Intn(upper)
	return randomValue
}

func RandomIndexValue(vector []int) int {
	i := rand.Intn(len(vector))
	return vector[i]
}
