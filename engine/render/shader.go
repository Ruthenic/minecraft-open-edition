package render

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
)

// Shader struct
type Shader struct {
	ID                       uint32
	vertexPath, fragmentPath string
}

// NewShader - create a new shader
func NewShader(vertexPath, fragmentPath string) (*Shader, error) {
	var s Shader
	vertexSource, err := LoadShader(vertexPath)
	if err != nil {
		return nil, err
	}
	fragmentSource, err := LoadShader(fragmentPath)
	if err != nil {
		return nil, err
	}
	vertexShader, err := CompileShader(vertexSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}
	fragmentShader, err := CompileShader(fragmentSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	s.ID = gl.CreateProgram()
	gl.AttachShader(s.ID, vertexShader)
	gl.AttachShader(s.ID, fragmentShader)
	gl.LinkProgram(s.ID)
	var status int32
	gl.GetProgramiv(s.ID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(s.ID, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(s.ID, logLength, nil, gl.Str(log))
		return nil, fmt.Errorf("failed to link program: %v", log)
	}
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	log.Printf("shader %v loaded successfully\n", s.ID)

	return &s, nil
}

// Use - opengl will use shader
func (s *Shader) Use() {
	gl.UseProgram(s.ID)
}

// LoadShader - loads shader program from source
func LoadShader(path string) (string, error) {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(source) + "\x00", nil
}

// CompileShader - compiles shader source
func CompileShader(source string, xtype uint32) (uint32, error) {
	shader := gl.CreateShader(xtype)
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}
	return shader, nil
}
