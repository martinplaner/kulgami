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
