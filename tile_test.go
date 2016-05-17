package main

import (
	"testing"
)

var t1 = Tile{
	x: 2,
	y: 0,
	w: 1,
	h: 3,
}

var t2 = Tile{
	x: 2,
	y: 1,
	w: 3,
	h: 2,
}

var t3 = Tile{
	x: 4,
	y: 2,
	w: 2,
	h: 2,
}

func TestOverlap(t *testing.T) {
	tests := []struct {
		t1      Tile
		t2      Tile
		overlap int
	}{
		{
			t1,
			t2,
			2,
		},
		{
			t1,
			t3,
			0,
		},
		{
			t2,
			t3,
			1,
		},
	}

	for _, test := range tests {
		o1 := test.t1.Overlap(test.t2)
		o2 := test.t2.Overlap(test.t1)

		if o1 != test.overlap {
			t.Errorf("Overlap(%v, %v) = %v, expected %v", test.t1, test.t2, o1, test.overlap)
		}

		if o1 != o2 {
			t.Errorf("Overlap(%v, %v) =/= Overlap(%v, %v)", test.t1, test.t2, test.t2, test.t1)
		}
	}
}
