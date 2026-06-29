package eagersingleton

type demo struct {
}

func newDemo() *demo {
	return &demo{}
}

func (d *demo) DoSomething() string {
	return "eager singleton demo is working..."
}

type Instance interface {
	DoSomething() string
}

var d *demo

func init() {
	d = newDemo()
}

func GetDemo() Instance {
	return d
}
