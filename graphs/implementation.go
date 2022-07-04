package graphs

import "log"

type Vertex struct {
	Name     string
	Adjacent *Vertex
}

type Graph struct {
	Vertices []*Vertex
}

//Inserts new vertex into graph
func (g *Graph) Insert(v string) {
	if g.HasKey(v) {
		log.Printf("Vertex %v already exists\n", v)
		return
	}
	newVertex := &Vertex{Name: v}
	g.Vertices = append(g.Vertices, newVertex)
}

//Checks if a key already exists in graph
func (g *Graph) HasKey(key string) bool {
	for _, v := range g.Vertices {
		if v.Name == key {
			return true
		}
	}
	return false
}

//Get vertext
func GetVertex(g *Graph, key string) *Vertex {
	for _, v := range g.Vertices {
		if v.Name == key {
			return v
		}
	}
	return nil
}

//Add friend
func (g *Graph) AddEdge(v1, v2 string) {
	if !g.HasKey(v1) || !g.HasKey(v2) {
		if !g.HasKey(v1) {
			log.Fatalf("Vertex %v does not exist\n", v1)
		} else {
			log.Fatalf("Vertex %v does not exist\n", v2)
		}
		return
	}
	for _, v := range g.Vertices {
		if v.Name == v1 {
			v.Adjacent = &Vertex{Name: v2}
		}
		if v.Name == v2 {
			v.Adjacent = &Vertex{Name: v1}
		}
	}
}

func (g *Graph) PrintValues() {
	for _, v := range g.Vertices {
		log.Printf("Name: %v, Friend: %v\n", v.Name, v.Adjacent)
	}
}
