// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"sort"
	"time"
)

const boardSize = 10

const popSize = 100
const generations = 2000

const mutateProb = 0.30

const fitChangeThreshold = 100
const fitChangeLength = 100

func main() {
	stopProfiling := initProfiling()

	rand.Seed(time.Now().UnixNano())

	pop := make([]Board, 0, popSize)

	for i := 0; i < popSize; i++ {
		pop = append(pop, RandomBoard())
	}

	fit := FitnessTracker{
		trackLength:        100,
		avgChangeThreshold: 5,
	}

	//	for gen := 0; gen < generations; gen++ {
	//	for SelectBestIndividual(pop).Fitness() < 0 {

	gen := 1
	for fit.IsImproving() {
		best10 := SelectByFitness(pop, 10)
		worthy := SelectByFitnessRoulette(pop, 20)
		children := CrossoverRandom(worthy, 70)
		Mutate(children)
		pop = append(worthy, children...)
		pop = append(pop, best10...)

		best := SelectBestIndividual(pop)
		fmt.Printf("Generation %04d: %d\n", gen, best.Fitness())
		fit.Add(best.Fitness())
		gen++
	}

	b := SelectBestIndividual(pop)
	fmt.Println(b, b.Fitness())
	exportBoard(b, "board.png")

	// open image
	open("board.png")

	stopProfiling()
}

func initProfiling() func() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	}

	return pprof.StopCPUProfile
}

// SelectByFitness returns the best n boards of the population,
// chosen by value of the fitness function. n must be greater than 0.
func SelectByFitness(pop []Board, n int) []Board {
	sort.Sort(sort.Reverse(ByFitness(pop)))
	return pop[0:n]
}

func SelectByFitnessRoulette(pop []Board, n int) []Board {
	sort.Sort(sort.Reverse(ByFitness(pop)))
	selected := make([]Board, 0, n)

	fitSum := 0
	leastFit := pop[len(pop)-1].Fitness()
	fits := make([]int, len(pop))

	if leastFit < 0 {
		leastFit *= -1
	} else {
		leastFit = 0
	}

	for i := 0; i < len(pop); i++ {
		fit := pop[i].Fitness() + leastFit + 1
		fitSum += fit
		fits[i] = fit
	}

	for len(selected) < n {
		v := rand.Intn(fitSum) + 1
		for j := 0; j < len(fits); j++ {
			v -= fits[j]
			if v <= 0 {
				selected = append(selected, pop[j])
				break
			}
		}
	}

	return selected
}

func SelectBestIndividual(pop []Board) Board {
	return SelectByFitness(pop, 1)[0]
}

func CrossoverPerm(pop []Board) []Board {
	perm := rand.Perm(len(pop))
	children := make([]Board, 0, len(pop))
	for i := 0; i < len(pop); i += 2 {
		c1, c2 := pop[perm[i]].Crossover(pop[perm[i+1]])
		children = append(children, c1, c2)
	}
	return children
}

func CrossoverRandom(pop []Board, n int) []Board {
	children := make([]Board, 0, n)
	for i := 0; i < n; i++ {
		// Choose parents
		p1 := rand.Intn(len(pop))
		p2 := rand.Intn(len(pop))

		c, _ := pop[p1].Crossover(pop[p2])
		children = append(children, c)
	}
	return children
}

func Mutate(pop []Board) {
	for i := 0; i < len(pop); i++ {
		if rand.Float32() < mutateProb {
			pop[i].Mutate()
		}
	}
}

//// SelectByFitness returns n boards of the population,
//// chosen by roulette wheel selection (ordered by value of the fitness function).
//func SelectByFitness(pop []Board, n int) []Board {
//	sort.Sort(sort.Reverse(ByFitness(pop)))
//	result = make([]Board, 0, n)

//	fitSum := 0
//	for _, b := range pop {
//		fitSum += b.Fitness()
//	}

//	for
//}

type ByFitness []Board

func (b ByFitness) Len() int {
	return len(b)
}

func (b ByFitness) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByFitness) Less(i, j int) bool {
	return b[i].Fitness() < b[j].Fitness()
}
