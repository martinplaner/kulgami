package main

import "math/rand"

const boundsPenality = 1000
const overlapPenalty = 1000
const sameNeighbourPenality = 10
const gapPenality = 10
const notConnectedPenalty = 5000

const neighbourCountBonus = 20

const rotateProb = 0.25
const translateProb = 0.15
const translateStdDev = 2
const swapPosProb = 0.1

/*
Board represents a single game board instance as an array of tiles.

Order of the tiles in the array is always:
- 1x2 (4x)
- 1x3 (4x)
- 2x2 (5x)
- 2x3 (4x)

i.e. [1x2 1x2 1x2 1x2 1x3 ... 2x2 2x3 2x3 2x3 2x3]

DO NOT CHANGE ORDER OF ARRAY ELEMENTS ONCE INTITIALIZED!
*/
type Board []Tile

func (b Board) Fitness() int {
	f := 0

	// Check tile bounds
	for _, t := range b {
		if t.x+t.w-1 >= boardSize {
			f -= boundsPenality
		}
		if t.y+t.h-1 >= boardSize {
			f -= boundsPenality
		}
	}

	// Check all tiles pair-wise
	for i := 0; i < len(b)-1; i++ {
		for j := i + 1; j < len(b); j++ {
			b1, b2 := b[i], b[j]
			f -= b1.Overlap(b2) * overlapPenalty
			f += b1.NeighbourCount(b2) * neighbourCountBonus

			if b1.Area() == b2.Area() {
				f -= sameNeighbourPenality * b1.Area()
			}

			f -= b1.GapSum(b2) * gapPenality
		}
	}

	if !b.IsConnected() {
		f -= notConnectedPenalty
	}

	return f
}

func (b Board) Crossover(other Board) (child1, child2 Board) {
	cutPoint := rand.Intn(tileNum-1) + 1

	child1 = append(child1, b[0:cutPoint]...)
	child1 = append(child1, other[cutPoint:len(other)]...)

	child2 = append(child2, other[0:cutPoint]...)
	child2 = append(child2, b[cutPoint:len(b)]...)

	return child1, child2
}

func (b Board) Mutate() {
	for i := 0; i < len(b); i++ {
		if rand.Float32() < rotateProb {
			b[i] = b[i].Rotate()
		}

		if rand.Float32() < translateProb {
			dx := int(rand.NormFloat64() * translateStdDev)
			dy := int(rand.NormFloat64() * translateStdDev)

			b[i].x = max(0, min(b[i].x+dx, boardSize-b[i].w))
			b[i].y = max(0, min(b[i].y+dy, boardSize-b[i].h))
		}

		if rand.Float32() < swapPosProb {
			j := rand.Intn(len(b))
			b[i].x, b[i].y = b[j].y, b[j].x
		}
	}
}

func (b Board) IsConnected() bool {
	visited := make(map[Tile]bool)
	q := new(Queue)

	q.Enqueue(b[0])

	for {
		t, ok := q.Dequeue()
		if !ok {
			break
		}

		if visited[t] {
			continue
		}
		visited[t] = true

		for i := 0; i < len(b); i++ {
			//			// TODO Is this form of equality enough?
			//			if b[i] == t {
			//				continue
			//			}

			if t.NeighbourCount(b[i]) > 0 {
				q.Enqueue(b[i])
			}
		}
	}

	return len(visited) == len(b)
}

// RandomIndividual generates a new random Board.
func RandomBoard() Board {
	b := Board{}

	var tileSpec = []struct {
		n int
		w int
		h int
	}{
		{4, 1, 2},
		{4, 1, 3},
		{5, 2, 2},
		{4, 2, 3},
	}

	for _, ts := range tileSpec {
		for i := 0; i < ts.n; i++ {
			b = append(b, Tile{
				x: rand.Intn(boardSize - ts.w + 1),
				y: rand.Intn(boardSize - ts.h + 1),
				w: ts.w,
				h: ts.h,
			})
		}
	}

	return b
}
