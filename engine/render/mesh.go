package render

import "github.com/go-gl/gl/v4.5-core/gl"

// TODO: refactor. not sure if my code is the best it can be

// Mesh struct
type Mesh struct {
	Vertices      []float32
	Indices       []uint32
	DrawMode, VAO uint32
}

// NewMesh - create a new mesh
func NewMesh(vertices []float32, indices []uint32, DrawMode uint32) *Mesh {
	var m Mesh
	m.Vertices = vertices
	m.Indices = indices
	m.DrawMode = DrawMode
	m.VAO = MakeVAO(vertices, indices, DrawMode)

	return &m
}

// MakeVAO - makes a vao from vertex data
func MakeVAO(vertices []float32, indices []uint32, drawMode uint32) uint32 {
	var vao, vbo, ebo uint32

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), drawMode)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(vertices), gl.Ptr(indices), drawMode)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.BindVertexArray(0)

	return vao
}

// Draw - draws the mesh with specified polygon mode
func (m *Mesh) Draw(mode uint32) {
	gl.PolygonMode(gl.FRONT_AND_BACK, mode)
	gl.BindVertexArray(m.VAO)
	gl.DrawElements(gl.TRIANGLES, int32(m.VertexCount()), gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.BindVertexArray(0)
}

// VertexCount - returns the number of vertices in the mesh
func (m *Mesh) VertexCount() int {
	return len(m.Indices)
}

// TriangleCount - returns the number of triangles in the mesh
func (m *Mesh) TriangleCount() int {
	return m.VertexCount() / 3
}

// TODO: write mesh drawing code
