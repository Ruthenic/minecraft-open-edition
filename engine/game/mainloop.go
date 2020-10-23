package game

import (
	"log"
	"runtime"

	"github.com/Broyojo/minecraft-go-edition/engine/render"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Setup - setup for game
func Setup() {
	log.Println("starting game...")
	runtime.LockOSThread() // OpenGL and GLFW must run on main os thread
}

// Run - runs the game
func Run() error {
	window, err := render.NewWindow(800, 600, "minecraft-go-edition")
	if err != nil {
		return err
	}
	defer glfw.Terminate()

	// main game loop
	for !window.ShouldClose() {
		// backround
		gl.ClearColor(1.0, 1.0, 1.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		/**game logic**/

		// plumbing
		glfw.PollEvents()
		window.SwapBuffers()
	}

	return nil
}
