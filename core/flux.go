package core

import (
	"fmt"
	"log"

	"github.com/nsf/termbox-go"
)

type FluxConfig struct {
    Foreground, Background termbox.Attribute
}

type Flux struct {
    Cursor Vec2
    Config FluxConfig
    Frame *Frame
    Buffer *Buffer
}

func NewFlux(config FluxConfig) *Flux {
    width, height := termbox.Size()
    flux := &Flux {
        Frame: NewFrame(0, 0, uint32(width), uint32(height)),
        Buffer: nil,
        Config: config,
        Cursor: Vec2Zeros(),
    }
    flux.Buffer = NewBuffer(&flux.Cursor)
    return flux
}

func NewDefaultFlux() *Flux {
    conf := FluxConfig{}
    conf.Background = termbox.ColorDefault
    conf.Foreground = termbox.ColorDefault
    return NewFlux(conf)
}

func (f *Flux) Start() {
    if err := termbox.Init(); err != nil {
        log.Fatal(err)
    }
    defer termbox.Close()

    running := true

    for running {
        switch ev := termbox.PollEvent(); ev.Type {
            case termbox.EventKey: 
            switch ev.Key {
            case termbox.KeyEsc:
                running = false
                break
            case termbox.KeyArrowUp:

            case termbox.KeyArrowDown:

            case termbox.KeyArrowLeft:

            case termbox.KeyArrowRight:

            case termbox.KeyEnter:
                f.Buffer.Write('\n')
            case termbox.KeyBackspace2:
                f.Buffer.Backspace()
            case termbox.KeyBackspace:
                f.Buffer.Backspace()
            default:
                f.Buffer.Write(byte(ev.Ch))
            }
        case termbox.EventError:
            log.Fatal(ev.Err)
        }
        if err := f.Render(); err != nil {
            log.Fatal(err)
        }
    }

    fmt.Println(f.Buffer)

}

func (f *Flux) Render() error {
    if err := termbox.Clear(f.Config.Background, f.Config.Foreground); err != nil {
        return err
    }

    f.Frame.RenderBuffer(f.Buffer, f.Config.Foreground, f.Config.Background)
    termbox.SetCursor(int(f.Cursor.X), int(f.Cursor.Y))
    return termbox.Flush()
}
