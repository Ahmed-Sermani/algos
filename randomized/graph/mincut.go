package graph

import (
	"math/rand"
	"time"
)

// Karger takes a graph pointer and find the min cut.
// Karger algorithm is randomized nondeterministic algorithms https://en.wikipedia.org/wiki/Karger%27s_algorithm .
// unlike the deterministic algorithms (max flow min cut) this algorithm takes O(n) where n
// is the number of edges.
func Karger(g *Graph) int {
	// to increase the probability of getting the mincut.
	// the algorithms will run multiable time and pick the min answer.
	min := len(g.Edges)
	for i := 0; i < 1000; i++ {
		if m := performKarger(g); m < min {
			min = m
		}
	}
	return min
}

func performKarger(g *Graph) int {
	// seed rand to make each run independent
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for g.Vertices > 2 {
		i := r.Intn(len(g.Edges))
		g.ContractEdge(i)
	}
	return len(g.Edges)
}
