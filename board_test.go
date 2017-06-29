// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	for i := 0; i < 1000; i++ {
		b := RandomBoard()
		if len(b) != tileNum {
			t.Errorf("number of tiles in new board was %v, expected %v", len(b), tileNum)
		}
	}
}

func TestCrossover(t *testing.T) {
	for i := 0; i < 1000; i++ {
		b := RandomBoard()
		c1, c2 := b.Crossover(b)
		if !reflect.DeepEqual(c1, c2) {
			t.Errorf("reflexive crossover not working!")
		}
	}
}

//func TestRandomBoardBounds(t *testing.T) {
//	for i := 0; i < 1000; i++ {
//		b := RandomBoard()
//		if b.Fitness() > 0 {
//			t.Errorf("fitness of new board was %v, expected %v", b.Fitness(), 0)
//		}
//	}
//}
