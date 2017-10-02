package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

var packagePath string

func init() {
	packagePath = filepath.Join(os.Getenv("GOPATH"), "src", "github.com/iand/mconeweek")
}

func loadShaders(vertexShader string, fragmentShader string) uint32 {

	vertexShaderID := compileShader(readShader(vertexShader), gl.VERTEX_SHADER)
	fragmentShaderID := compileShader(readShader(fragmentShader), gl.FRAGMENT_SHADER)

	shaderID := linkProgram(vertexShaderID, fragmentShaderID)

	gl.DeleteShader(vertexShaderID)
	gl.DeleteShader(fragmentShaderID)

	return shaderID
}

func compileShader(source string, shaderType uint32) uint32 {
	shaderID := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shaderID, 1, csources, nil)
	free()

	gl.CompileShader(shaderID)

	var status int32
	gl.GetShaderiv(shaderID, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderID, logLength, nil, gl.Str(log))

		panic(fmt.Sprintf("failed to compile %v: %v", source, log))
	}

	return shaderID

}

func linkProgram(vertexShaderID uint32, fragmentShaderID uint32) uint32 {
	id := gl.CreateProgram()

	gl.AttachShader(id, vertexShaderID)
	gl.AttachShader(id, fragmentShaderID)

	gl.LinkProgram(id)

	return id
}

func readShader(name string) string {
	filename := filepath.Join(packagePath, "Shaders", name)
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("shader file not found at %s", filename))
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(fmt.Sprintf("failed to read shader file at %s", filename))
	}

	return string(b)
}
