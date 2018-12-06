package model

import "testing"

func TestFloat64Round(t *testing.T) {
	t.Log(3.13131313214432, Float64Round(3.13131313214432, 2))
}
