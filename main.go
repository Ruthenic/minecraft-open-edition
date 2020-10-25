package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/Broyojo/minecraft-open-edition/engine/game"
	"github.com/Broyojo/minecraft-open-edition/engine/voxel"
)

func init() {
	runtime.LockOSThread() // OpenGL and GLFW must run on main os thread
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := voxel.NewChunk(1, 2, 3)
	randomize(c)

	fmt.Println(c.Blocks[4][5][6])

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}

func randomize(c *voxel.Chunk) {
	c.Map(func(x, y, z int) voxel.Block {
		return voxel.Block(rand.Intn(voxel.NumBlocks))
	})
}
