package functions

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"time"
)

func NewRandomValue(upper int) int {
	randomValue := rand.Intn(upper)
	return randomValue
}

func RandomIndexValue(vector []int) int {
	i := rand.Intn(len(vector))
	return vector[i]
}

func RandomPhoneWithId(id int) string {
	rand.NewSource(time.Now().UnixNano())
	phoneNumber := strconv.Itoa(id)

	for i := 0; i < 9; i++ {
		digit := rand.Intn(10)
		phoneNumber += fmt.Sprintf("%d", digit)
	}
	return phoneNumber
}

func RandomAddress() string {
	return uuid.New().String()
}
