package graph

// Vertex represents a vertex in a graph.
type Vertex struct {
	id         int64
	properties map[string]interface{}
}

// NewVertex creates a new vertex.
func NewVertex(id int64) (*Vertex, error) {
	vertex := &Vertex{
		id:         id,
		properties: make(map[string]interface{}),
	}
	return vertex, nil
}
