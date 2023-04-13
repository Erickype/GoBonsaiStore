package functions

import (
	"math/rand"
	"time"
)

func GenerateDateInRange(lower int, upper int) string {
	start := time.Date(lower, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(upper, 1, 1, 0, 0, 0, 0, time.UTC)

	randomSeconds := rand.Int63n(end.Unix()-start.Unix()) + start.Unix()
	randomTime := time.Unix(randomSeconds, 0)

	return randomTime.Format("2006-01-02 15:04:05.999999-07:00")
}
