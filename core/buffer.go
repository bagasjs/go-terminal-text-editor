package core

import (
	"fmt"
	"strings"
)

type Buffer struct {
    Cursor *Vec2
    lines []string
}

func NewBuffer(cursor *Vec2) *Buffer {
    return &Buffer{
        lines: make([]string, 1),
        Cursor: cursor,
    }
}

func (b *Buffer) InsertEmptyLine(index uint64) {
    b.lines = append(b.lines, "")
    bufferLinesAmount := len(b.lines)
    if index > uint64(bufferLinesAmount) {
        index = uint64(bufferLinesAmount)
    }

    for i := index + 1; i < uint64(bufferLinesAmount); i++ {
        b.lines[i] = b.lines[i - 1]
    }

    b.lines[index] = ""
}

func (b *Buffer) Write(ch byte) {
    pointedLine := b.lines[b.Cursor.Y]
    if ch == '\n' {
        removedPart := pointedLine[b.Cursor.X:]
        leftPart := pointedLine[:b.Cursor.X]
        b.lines[b.Cursor.Y] = leftPart
        b.Cursor.Y += 1
        b.InsertEmptyLine(uint64(b.Cursor.Y))
        b.lines[b.Cursor.Y] = removedPart
        b.Cursor.X = 0
    } else if ch == '\r' {
        return
    } else {
        b.lines[b.Cursor.Y] = fmt.Sprint(pointedLine[:b.Cursor.X], string(ch), pointedLine[b.Cursor.X:])
        b.Cursor.X += 1
    }
}

func (b* Buffer) RemoveLine(index uint64) string {
    removedLine := b.lines[index]

    bufferLinesAmount := len(b.lines)
    if index >= uint64(bufferLinesAmount) {
        index = uint64(bufferLinesAmount) - 1
    }

    for i := index; i < uint64(bufferLinesAmount) - 1; i++ {
        b.lines[i - 1] = b.lines[i + 1]
    }

    b.lines = b.lines[:bufferLinesAmount-1]

    return removedLine
}


func (b* Buffer) WriteText(text string) {
    for _, ch := range text {
        b.Write(byte(ch))
    }
}

func (b* Buffer) Backspace() byte {
    if b.Cursor.X == 0 {
        if b.Cursor.Y != 0 {
            removedLine := b.RemoveLine(uint64(b.Cursor.Y))
            b.Cursor.Y -= 1
            b.Cursor.X = int32(len(b.lines[b.Cursor.Y]))
            b.lines[b.Cursor.Y] = fmt.Sprint(b.lines[b.Cursor.Y], removedLine)
            return '\n'
        }
    } else {
        b.Cursor.X -= 1
        modifiedLine := b.lines[b.Cursor.Y]
        removedChar := modifiedLine[b.Cursor.X]
        b.lines[b.Cursor.Y] = fmt.Sprint(modifiedLine[:b.Cursor.X], modifiedLine[b.Cursor.X+1:])
        return removedChar
    }

    return 0;
}

func (b *Buffer) ToString() string {
    result := strings.Join(b.lines, "\n")
    return result
}

