package abstractFF

import "fmt"

type GoodFarmerFactory struct{}

func (g *GoodFarmerFactory) myAspect() {
	fmt.Println("good farmer aspect...")
}

func (g *GoodFarmerFactory) CreateStrawberry() Strawberry {
	g.myAspect()
	defer g.myAspect()
	return &GoodFarmerStrawberry{
		brand: "goodfarmer",
	}
}

func (g *GoodFarmerFactory) CreateLemon() Lemon {
	g.myAspect()
	defer g.myAspect()
	return &GoodFarmerLemon{
		brand: "goodfarmer",
	}
}

type DoleFactory struct{}

func (d *DoleFactory) myAspect() {
	fmt.Println("dole aspect...")
}

func (d *DoleFactory) CreateStrawberry() Strawberry {
	d.myAspect()
	defer d.myAspect()
	return &DoleStrawberry{
		brand: "dole",
	}
}

func (d *DoleFactory) CreateLemon() Lemon {
	d.myAspect()
	defer d.myAspect()
	return &DoleLemon{
		brand: "dole",
	}
}
