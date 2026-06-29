package factory

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type fruitCreator func(name string) Fruit

type SimpleFruitFactory struct {
	creator map[string]fruitCreator
}

func NewSimpleFF() *SimpleFruitFactory {
	return &SimpleFruitFactory{
		creator: map[string]fruitCreator{
			"orange":     NewOrange,
			"strawberry": NewStrawberry,
			"cherry":     NewCherry,
		},
	}
}

func (sf *SimpleFruitFactory) CreateFruit(typ string) (Fruit, error) {
	fruitCreator, ok := sf.creator[typ]
	if !ok {
		return nil, fmt.Errorf("fruit typ: %s is not supported yet.", typ)
	}

	src := rand.NewSource(time.Now().UnixNano())
	rander := rand.New(src)
	name := strconv.Itoa(rander.Int())
	return fruitCreator(name), nil
}
