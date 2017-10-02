package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Model struct {
	vao          uint32
	vboCount     uint32
	indicesCount int32
	buffers      []uint32
}

func NewModel(vertexPositions []float32, textureCoords []float32, indices []uint32) *Model {
	m := Model{
		buffers: []uint32{},
	}
	m.addData(vertexPositions, textureCoords, indices)
	return &m
}

func (m *Model) dispose() {
	m.deleteData()
}

func (m *Model) bindVAO() {
	gl.BindVertexArray(m.vao)
}

func (m *Model) addData(vertexPositions []float32, textureCoords []float32, indices []uint32) {
	if m.vao != 0 {
		m.deleteData()
	}

	m.indicesCount = int32(len(indices))

	gl.GenVertexArrays(1, &m.vao)
	gl.BindVertexArray(m.vao)

	m.addVBO(3, vertexPositions)
	m.addVBO(2, textureCoords)
	m.addEBO(indices)
}

func (m *Model) addVBO(dimensions int32, data []float32) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)

	gl.VertexAttribPointer(m.vboCount, dimensions, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(m.vboCount)
	m.vboCount++

	m.buffers = append(m.buffers, vbo)
}

func (m *Model) addEBO(indices []uint32) {
	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
}

func (m *Model) deleteData() {
	gl.DeleteVertexArrays(1, &m.vao)
	gl.DeleteBuffers(int32(len(m.buffers)), &m.buffers[0])

	m.buffers = m.buffers[:0]

	m.vboCount = 0
	m.vao = 0
	m.indicesCount = 0
}

func (m *Model) getIndicesCount() int32 {
	return m.indicesCount
}
