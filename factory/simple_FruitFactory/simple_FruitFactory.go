package simpleFF

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/tilosrHSZ/go-design-patterns/factory/fruit"
)

type fruitCreator func(name string) fruit.Fruit

type SimpleFruitFactory struct {
	creator map[string]fruitCreator
}

func NewSimpleFF() *SimpleFruitFactory {
	return &SimpleFruitFactory{
		creator: map[string]fruitCreator{
			"orange":     fruit.NewOrange,
			"strawberry": fruit.NewStrawberry,
			"cherry":     fruit.NewCherry,
		},
	}
}

func (sf *SimpleFruitFactory) CreateFruit(typ string) (fruit.Fruit, error) {
	fruitCreator, ok := sf.creator[typ]
	if !ok {
		return nil, fmt.Errorf("fruit typ: %s is not supported yet.", typ)
	}

	src := rand.NewSource(time.Now().UnixNano())
	rander := rand.New(src)
	name := strconv.Itoa(rander.Int())
	return fruitCreator(name), nil
}
