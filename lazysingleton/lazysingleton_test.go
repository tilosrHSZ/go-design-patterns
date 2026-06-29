package lazysingleton

import (
	"testing"
)

func Test_singleton2(t *testing.T) {
	d := GetDemo()
	t.Error(d.Work())
}
