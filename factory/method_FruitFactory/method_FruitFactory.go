package methodFF

import "github.com/tilosrHSZ/go-design-patterns/factory/fruit"

type FruitFactory interface {
	CreateFruit() fruit.Fruit
}

type OrangeFactory struct {
}

func NewOrangeFactory() FruitFactory {
	return &OrangeFactory{}
}

func (o *OrangeFactory) CreateFruit() fruit.Fruit {
	return fruit.NewOrange("")
}

type StrawberryFactory struct {
}

func NewStrawberryFactory() FruitFactory {
	return &StrawberryFactory{}
}

func (s *StrawberryFactory) CreateFruit() fruit.Fruit {
	return fruit.NewStrawberry("")
}

type CherryFactory struct {
}

func NewCherryFactory() FruitFactory {
	return &CherryFactory{}
}

func (c *CherryFactory) CreateFruit() fruit.Fruit {
	return fruit.NewCherry("")
}
