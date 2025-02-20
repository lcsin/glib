package pkg

import "testing"

func TestBool2Int(t *testing.T) {
	t.Log(Bool2Int(true))
}

func TestNanoTime(t *testing.T) {
	t.Log(NanoTime())
}
