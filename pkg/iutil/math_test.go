package iutil

import "testing"

func TestCeilInts(t *testing.T) {
	t.Log(CeilInts(4, 2))
	t.Log(CeilInts(5, 2))
}
