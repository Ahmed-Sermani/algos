package graph_test

import (
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/Ahmed-Sermani/algos/graph"
)

func graphFromFile(name string) *graph.Graph {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("error reading " + name)
	}
	g := graph.New()

	// parsing
	for _, line := range strings.Split(string(data), "\n") {
		columns := strings.Split(line, "\t")
		tail, _ := strconv.Atoi(columns[0])
		for _, col := range columns[1:] {
			head, _ := strconv.Atoi(col)
			g.AddEdge(tail, head)
		}
		g.Vertices++
	}
	return g
}

func TestMinCutShort(t *testing.T) {
	g := graph.New()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	g.AddEdge(1, 0)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 0)
	g.AddEdge(3, 1)
	g.AddEdge(3, 2)
	g.Vertices = 4
	t.Log(graph.Karger(g))
}

func TestMinCutLong(t *testing.T) {
	g := graphFromFile("testdata/inputLong.txt")
	graph.Karger(g)
}
