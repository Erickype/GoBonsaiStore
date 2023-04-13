package functions

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomEmailAddress(firstname, lastname string) string {
	var email string
	rand.NewSource(time.Now().UnixNano())
	number := strconv.Itoa(NewRandomValue(99))
	email = firstname[0:2] + lastname[0:2] + number + RandomIndexString(emailProviders)
	return email
}
