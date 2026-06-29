package test

import (
	"testing"

	"github.com/tilosrHSZ/go-design-patterns/lazysingleton"
)

func Test(t *testing.T) {
	l := lazysingleton.GetDemo()
	t.Error(l.Work())
}
