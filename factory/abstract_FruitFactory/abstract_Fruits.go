package abstractFF

import "fmt"

type Strawberry interface {
	SweetAttack()
}

type Lemon interface {
	AcidAttack()
}

type FruitFactory interface {
	CreateStrawberry() Strawberry
	CreateLemon() Lemon
}

type GoodFarmerStrawberry struct {
	brand string
	Strawberry
}

func (g *GoodFarmerStrawberry) SweetAttack() {
	fmt.Printf("sweet attack from %s", g.brand)
}

type GoodFarmerLemon struct {
	brand string
	Lemon
}

func (g *GoodFarmerLemon) AcidAttack() {
	fmt.Printf("acid attack from %s", g.brand)
}

type DoleStrawberry struct {
	brand string
	Strawberry
}

func (d *DoleStrawberry) SweetAttack() {
	fmt.Printf("sweet attack from %s", d.brand)
}

type DoleLemon struct {
	brand string
	Lemon
}

func (d *DoleLemon) AcidAttack() {
	fmt.Printf("acid attack from %s", d.brand)
}
