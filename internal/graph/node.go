package graph

// Node represents a node in a graph.
type Node struct {
	id         int64
	properties map[string]interface{}
}

// NewNode creates a new node.
func NewNode(id int64) (*Node, error) {
	vertex := &Node{
		id:         id,
		properties: make(map[string]interface{}),
	}
	return vertex, nil
}
