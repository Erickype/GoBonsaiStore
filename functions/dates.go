package functions

import (
	"math/rand"
	"time"
)

func GenerateDateRange() time.Time {
	start := time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	randomSeconds := rand.Int63n(end.Unix()-start.Unix()) + start.Unix()
	randomTime := time.Unix(randomSeconds, 0)

	return randomTime
}
