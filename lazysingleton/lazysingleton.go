package lazysingleton

import (
	"sync"
)

type demo struct {
}

func newDemo() *demo {
	return &demo{}
}

func (d *demo) Work() string {
	return "demo is working..."
}

type Instance interface {
	Work() string
}

var d *demo
var once sync.Once

func GetDemo() Instance {
	once.Do(func() {
		d = newDemo()
	})
	return d
}
