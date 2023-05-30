package functions

import (
	"github.com/goombaio/namegenerator"
	"math/rand"
	"strings"
	"time"
)

func RandomPersonName() (string, string, string) {
	rand.NewSource(time.Now().UnixNano())

	var sex = "M"
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	if strings.LastIndex(firstName, "a") == len(firstName)-1 {
		sex = "F"
	}

	return firstName, lastName, sex
}

func RandomFancyName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	return name
}

func RandomProfessionSalary() (string, float64) {
	profession := professions[rand.Intn(len(professions))]
	return profession.name, profession.salary
}
