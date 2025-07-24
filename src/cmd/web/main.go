//go:build js

package main

import (
	"syscall/js"

	mbset "github.com/alan-b-lima/mandelbrot-set/src/internal/mandelbrot_set"
)

var (
	_ImageData        = js.Global().Get("ImageData")
	createImageBitmap = js.Global().Get("window").Get("createImageBitmap")
)

func generateMandelbrot(_ js.Value, args []js.Value) any {
	set := mbset.MandelbrodSet{
		Origin: complex(args[0].Float(), args[1].Float()),
		Scale:  args[2].Float(),
		Width:  args[3].Int(), Height: args[4].Int(),
		IterationLimit: args[5].Int(),
	}

	set.Generate()

	arr := js.Global().Get("Uint8ClampedArray").New(len(set.Image))
	js.CopyBytesToJS(arr, set.Image)

	return js.ValueOf(map[string]any{
		"buffer": arr,
		"width":  set.Width,
		"height": set.Height,
	})
}

func main() {
	js.Global().Set("generateMandelbrot", js.FuncOf(generateMandelbrot))

	<-make(chan struct{})
}
