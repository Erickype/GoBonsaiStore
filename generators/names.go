package generators

import (
	"github.com/goombaio/namegenerator"
	"time"
)

func GenerateRandomName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	return name
}
