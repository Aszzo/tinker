package main

import (
	"github.com/Aszzo/tinker/file"
	"testing"
)

func TestResizeJpg(t *testing.T) {
	file.ResizeJpg("test/demo1/test.jpg", 70)
}
func TestResizePng(t *testing.T)  {
	file.ResizePng("test/test.png", 70, "test")
}