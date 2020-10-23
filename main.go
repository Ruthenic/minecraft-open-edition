package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Broyojo/minecraft-go-edition/engine/game"
	"github.com/Broyojo/minecraft-go-edition/engine/voxel"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	c := voxel.NewChunk(1, 2, 3)
	randomize(c)

	fmt.Println(c.Blocks[4][5][6])

	game.Setup()
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}

func randomize(c *voxel.Chunk) {
	c.Map(func(x, y, z int) voxel.Block {
		return voxel.Block(rand.Intn(voxel.NumBlocks))
	})
}
