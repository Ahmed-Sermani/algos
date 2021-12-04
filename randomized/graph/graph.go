package graph

type Graph struct {
	Vertices int
	Edges    []edge
}

type edge struct {
	tail int
	head int
}

func New() *Graph {
	return &Graph{}
}

func (g *Graph) AddEdge(tail, head int) {
	e := edge{tail, head}
	g.Edges = append(g.Edges, e)
}

func (g *Graph) AddVertex(label int) {
	g.Vertices++
}

// ContractEdge will collapse an edge and merge vertex elements
func (g *Graph) ContractEdge(i int) {
	e := g.Edges[i]

	// update edges to remove e
	g.Edges = append(g.Edges[:i], g.Edges[i+1:]...)

	// loop to remove e.head from edges slice
	for j, edge := range g.Edges {
		if edge.head == e.head {
			g.Edges[j].head = e.tail
		} else if edge.tail == e.head {
			g.Edges[j].tail = e.tail
		}
	}

	// remove self loops
	temp := []edge{}
	for _, edge := range g.Edges {
		if edge.head != edge.tail {
			temp = append(temp, edge)
		}
	}
	g.Edges = temp

	g.Vertices--
}
