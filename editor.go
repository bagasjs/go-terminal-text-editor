package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/nsf/termbox-go"
)

type Buffer struct {
    lines []string
    cursorPosX int
    cursorPosY int
}

func NewBuffer() *Buffer {
    return &Buffer{
        lines: make([]string, 1),
        cursorPosX: 0,
        cursorPosY: 0,
    }
}

func (b* Buffer) GetCurrentLine() int {
    return b.cursorPosY
}

func (b* Buffer) NextLine() {
    b.cursorPosY += 1
    b.cursorPosX = 0
    b.lines = append(b.lines, "")
}

func (b *Buffer) Write(text string) {
    text = strings.ReplaceAll(text, "\r\n", "\n")
    lines := strings.Split(text, "\n")
    b.lines[b.cursorPosY] = fmt.Sprintf("%s%s", b.lines[b.cursorPosY], lines[0])
    lines = lines[1:]
    b.lines = append(b.lines, lines...)
    b.UpdateCursorPosition()
}

func (b* Buffer) UpdateCursorPosition() {
    b.cursorPosX = len(b.lines[len(b.lines)-1])
}

func (b* Buffer) Backspace() {
    currentLine := b.lines[b.cursorPosY]
    currentLineLen := len(currentLine)
    if currentLineLen == 0 {
        if b.cursorPosY != 0 {
            b.lines = b.lines[:len(b.lines) - 1]
            b.cursorPosY -= 1
       }
    } else {
        b.lines[b.cursorPosY] = currentLine[:currentLineLen - 1]
    }
    b.UpdateCursorPosition()
}

func (b* Buffer) SaveToFile(filePath string) {
    err := ioutil.WriteFile(filePath, []byte(b.ToString()), 0664)
    if err != nil {
        log.Fatal(err)
    }
}

func (b *Buffer) Render(x, y int, fg, bg termbox.Attribute) {
    for i, line := range b.lines {
        for j, c := range line {
            termbox.SetCell(x + j, y + i, c, fg, bg)
        }
    }
    termbox.SetCursor(b.cursorPosX, b.cursorPosY)
}

func (b *Buffer) ToString() string {
    result := strings.Join(b.lines, "\n")
    result = strings.TrimPrefix(result, " ")
    return result
}

type Editor struct {
    Foreground termbox.Attribute
    Background termbox.Attribute

    buffer *Buffer
}

func NewEditor() *Editor {
    return &Editor{
        Foreground: termbox.ColorDefault,
        Background: termbox.ColorDefault,
        buffer: NewBuffer(),
    }
}

func (e *Editor) Start() {
    if err := termbox.Init(); err != nil {
        log.Fatal(err)
    }
    defer termbox.Close()

    for {
        switch ev := termbox.PollEvent(); ev.Type {
        case termbox.EventKey: 
            switch ev.Key {
            case termbox.KeyEsc:
                return
            case termbox.KeyEnter:
                e.buffer.NextLine()
            case termbox.KeyBackspace2:
                e.buffer.Backspace()
            case termbox.KeyBackspace:
                e.buffer.Backspace()
            case termbox.KeyCtrlS:
                e.buffer.SaveToFile("result.txt")
            default:
                e.buffer.Write(string(ev.Ch))
            }
        case termbox.EventError:
            log.Fatal(ev.Err)
        }
        if err := e.Render(); err != nil {
            log.Fatal(err)
        }
    }
}

func (e *Editor) Render() error {
    if err := termbox.Clear(e.Background, e.Foreground); err != nil {
        return err;
    }
    e.buffer.Render(0, 0, e.Foreground, e.Background) 
    return termbox.Flush()
}
