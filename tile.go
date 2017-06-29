// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

const tileNum = 17

type Tile struct {
	x int
	y int
	w int
	h int
}

func (t Tile) Rotate() Tile {
	return Tile{
		x: t.x,
		y: t.y,
		w: t.h,
		h: t.w,
	}
}

func (t Tile) Overlap(other Tile) int {
	xo := max(0, min(t.x+t.w, other.x+other.w)-max(t.x, other.x))
	yo := max(0, min(t.y+t.h, other.y+other.h)-max(t.y, other.y))

	return xo * yo
}

func (t Tile) GapSum(other Tile) int {
	xo := max(0, (min(t.x+t.w, other.x+other.w)-max(t.x, other.x))*-1)
	yo := max(0, (min(t.y+t.h, other.y+other.h)-max(t.y, other.y))*-1)

	return xo + yo
}

func (t Tile) NeighbourCount(other Tile) int {
	hExpand := Tile{
		x: max(0, t.x-1),
		y: t.y,
		w: t.w + 1,
		h: t.h,
	}

	vExpand := Tile{
		x: t.x,
		y: max(0, t.y-1),
		w: t.w,
		h: t.h + 1,
	}

	return hExpand.Overlap(other) + vExpand.Overlap(other)
}

func (t Tile) Area() int {
	return t.w * t.h
}
