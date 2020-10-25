package render

// Mesh struct
type Mesh struct {
	Vertices []Vertex
	Indices  []uint32
	Textures []Texture
}

// NewMesh - create a new mesh
func NewMesh(vertices []Vertex, indices []uint32, textures []Texture) *Mesh {
	var m Mesh
	m.Vertices = vertices
	m.Indices = indices
	m.Textures = textures
	return &m
}

// TODO: write mesh drawing code
