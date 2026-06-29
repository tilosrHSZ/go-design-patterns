package test

import (
	"testing"

	"github.com/tilosrHSZ/go-design-patterns/eagersingleton"
	"github.com/tilosrHSZ/go-design-patterns/lazysingleton"
)

func Test(t *testing.T) {
	e := eagersingleton.GetDemo()
	t.Error(e.DoSomething())

	l := lazysingleton.GetDemo()
	t.Error(l.Work())

}
