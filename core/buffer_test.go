package core_test

import (
	"fmt"

	"github.com/bagasjs/flux/core"
)

func testBufferWrite() {
    frame := core.NewFrame(0, 0, 100, 100)
    buffer := core.NewBuffer(&frame.Cursor)

    buffer.WriteText("Hello, World\nThis is Bagas Jonathan Sitanggang");

    frame.Cursor.X = 7
    frame.Cursor.Y = 0

    buffer.WriteText("World ")

    fmt.Println(buffer.ToString())
}

func testBufferBackspace() {
    frame := core.NewFrame(0, 0, 100, 100)
    buffer := core.NewBuffer(&frame.Cursor)

    buffer.WriteText("Hello, World\nThis is Bagas Jonathan Sitanggang");

    frame.Cursor.X = 7
    frame.Cursor.Y = 0

    for i := 0; i < 5; i++ {
        buffer.Backspace()
    }

    fmt.Println(buffer.ToString())
    fmt.Printf("Cursor(%d, %d)", frame.Cursor.X, frame.Cursor.Y)
}

