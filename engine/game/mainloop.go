package game

import (
	"log"

	"github.com/Broyojo/minecraft-open-edition/engine/render"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Run - runs the game
func Run() error {
	window, err := render.NewWindow(800, 600, "minecraft-go-edition")
	if err != nil {
		return err
	}
	defer glfw.Terminate()

	shader, err := render.NewShader("assets/shaders/vertex.vert", "assets/shaders/fragment.frag")
	if err != nil {
		return err
	}

	vertices := []float32{
		0.5, 0.5, 0,
		0.5, -0.5, 0,
		-0.5, -0.5, 0,
		-0.5, 0.5, 0,
	}

	indices := []uint32{
		0, 1, 3,
		1, 2, 3,
	}

	mesh := render.NewMesh(vertices, indices, gl.STATIC_DRAW)

	log.Printf("Loaded mesh %v with %v vertices and %v triangles", mesh.VAO, mesh.VertexCount(), mesh.TriangleCount())

	// main game loop
	for !window.ShouldClose() {
		// backround
		gl.ClearColor(0.5, 0.5, 0.5, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		/**game logic**/
		shader.Use()
		mesh.Draw(gl.FILL)
		// mesh.Draw(gl.LINE)

		// plumbing
		glfw.PollEvents()
		window.SwapBuffers()
	}

	return nil
}
