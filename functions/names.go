package functions

import (
	"github.com/goombaio/namegenerator"
	"math/rand"
	"strings"
	"time"
)

func RandomPersonName() (string, string, string) {
	rand.NewSource(time.Now().UnixNano())

	var sex = "male"
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	if strings.LastIndex(firstName, "a") == len(firstName)-1 {
		sex = "female"
	}

	return firstName, lastName, sex
}

func RandomFancyName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	return name
}
