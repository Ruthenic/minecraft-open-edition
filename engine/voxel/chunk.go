package voxel

// ChunkSize - size of chunk on all sides
const ChunkSize = 16

// Chunk struct
type Chunk struct {
	X, Y, Z int
	Blocks  [ChunkSize][ChunkSize][ChunkSize]Block
}

// NewChunk - creates a new chunk pointer
func NewChunk(x, y, z int) *Chunk {
	var c Chunk
	c.X, c.Y, c.Z = x, y, z
	return &c
}

// Map - maps a function across all blocks
func (c *Chunk) Map(f func(x, y, z int) Block) {
	for x := 0; x < ChunkSize; x++ {
		for y := 0; y < ChunkSize; y++ {
			for z := 0; z < ChunkSize; z++ {
				c.Blocks[x][y][z] = f(x, y, z)
			}
		}
	}
}
