package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type BasicTexture struct {
	id uint32
}

func NewBasicTexture(name string) *BasicTexture {
	b := BasicTexture{}
	b.loadFromFile(name)
	return &b
}

func (b *BasicTexture) loadFromFile(name string) {
	filename := filepath.Join(packagePath, "Res", "Textures", name+".png")
	imgFile, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("texture file not found at %s", filename))
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(fmt.Sprintf("failed to decode texture file at %s", filename))
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic(fmt.Sprintf("unsupported stride in texture, got %d, wanted %d", rgba.Stride, rgba.Rect.Size().X*4))
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	gl.GenTextures(1, &b.id)
	gl.BindTexture(gl.TEXTURE_2D, b.id)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix),
	)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
}

func (b *BasicTexture) bindTexture() {
	gl.BindTexture(gl.TEXTURE_2D, b.id)
}

func (b *BasicTexture) dispose() {
	gl.DeleteTextures(1, &b.id)
}
