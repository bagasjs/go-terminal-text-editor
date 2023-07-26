package core

import "github.com/nsf/termbox-go"

type Frame struct {
    Vec2
    Width, Height uint32
}

func NewFrame(x, y int32, width, height uint32) *Frame {
    return &Frame {
        Width: width,
        Height: height,
        Vec2: Vec2{ X: x, Y: y },
    }
}

func (f *Frame) RenderBuffer(buffer* Buffer, fg, bg termbox.Attribute) {
    for i, line := range buffer.lines {
        for j, c := range line {
            termbox.SetCell(int(f.X) + j, int(f.Y) + i, c, fg, bg)
        }
    }
}
