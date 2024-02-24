package main 

type Graph struct {
	vertices map[string]map[string]bool
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]map[string]bool)
	}
}

func (g *Graph)AddVertex(vertex string) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = make(map[string]bool)
	}
}

func(g *Graph)RemoveVertex(vertex string) {
	if _, exists := g.vertices[vertex]; exists {
		for neighbour := range g.vertices[vertex] {
			delete(g.vertices[neighbour], vertex)
		}
		delete(g.vertices[vertex])
	}
}

func (g *Graph)AddEdge(vertex1 string, vertex2 string) {
	g.AddVertex(vertex1)
	g.AddVertex(vertex2)
	g.vertices[vertex1][vertex2]=true
	g.vertices[vertex2][vertex1]=true
}

func (g *Graph)RemoveEdge(vertex1, vertex2 string){
	if _, exists1 := g.vertices[vertex1]; exists1 {
		if _, exists2 := g.vertices[vertex2]; exists2 {
			delete(g.vertices[vertex1], vertex2)
			delete(g.vertices[vertex2], vertex1)
		}
	}
}

func main() {
	
}