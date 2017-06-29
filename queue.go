// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Queue struct {
	elems []Tile
}

func (q *Queue) Enqueue(t Tile) {
	q.elems = append(q.elems, t)
}

func (q *Queue) Dequeue() (t Tile, ok bool) {
	if len(q.elems) < 1 {
		return t, false
	}

	elem := q.elems[0]
	q.elems = q.elems[1:len(q.elems)]

	return elem, true
}

func (q *Queue) Values() []Tile {
	return q.elems
}
