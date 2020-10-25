package game

import (
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

	// main game loop
	for !window.ShouldClose() {
		// backround
		gl.ClearColor(1.0, 1.0, 1.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		/**game logic**/
		shader.Use()

		// plumbing
		glfw.PollEvents()
		window.SwapBuffers()
	}

	return nil
}
