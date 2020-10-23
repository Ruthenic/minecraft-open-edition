package render

import (
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// NewWindow - returns a new glfw window
func NewWindow(width, height int, name string) (*glfw.Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	window, err := glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		return nil, err
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version:", version)
	return window, nil
}
