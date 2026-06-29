package singleton

import (
	"testing"
)

func Test_singleton(t *testing.T) {
	d := GetDemo()
	putInstance(d)
	t.Error(d.DoSomething())
}

func putInstance(d Instance) {

}
