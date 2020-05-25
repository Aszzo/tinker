package main

import (
	"github.com/img-modify/file"
	"testing"
)

func TestResizeJpg(t *testing.T) {
	file.ResizeJpg("test1.jpg")
}
func TestResizePng(t *testing.T)  {
	file.ResizePng("test/test.png")
}