package voxel

// Block Type Alias
type Block int

// Block Enum
const (
	Air Block = iota
	Dirt
	Stone
	NumBlocks int = iota
)
