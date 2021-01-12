package graph

// Edge represents the edge between two vertices in a graph.
type Edge struct {
	id         int64
	inboundID  int64
	outboundID int64
	properties map[string]interface{}
}

// NewEdge creates a new edge.
func NewEdge(id int64, inboundID int64, outboundID int64) (*Edge, error) {
	edge := &Edge{
		id:         id,
		inboundID:  inboundID,
		outboundID: outboundID,
		properties: make(map[string]interface{}),
	}
	return edge, nil
}
