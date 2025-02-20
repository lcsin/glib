package adapter

import "testing"

func TestTypeCAdapterUSB(t *testing.T) {
	expect := "usb phone charge"
	usb := NewUSB()
	typeC := NewTypeCAdapter(usb)
	res := typeC.TypeCCharge()
	if expect != res {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
