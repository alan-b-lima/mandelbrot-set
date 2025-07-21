//go:build js

package main

import (
	"syscall/js"

	mbset "github.com/alan-b-lima/mandelbrot-set/src/internal/multibrot_set"
)

var (
	_ImageData        = js.Global().Get("ImageData")
	createImageBitmap = js.Global().Get("window").Get("createImageBitmap")
)

func generateMandelbrot(_ js.Value, args []js.Value) any {
	set := mbset.MultibrodSet{
		Seed:   complex(args[0].Float(), args[1].Float()),
		Power:  complex(args[2].Float(), args[3].Float()),
		Origin: complex(args[4].Float(), args[5].Float()),
		Scale:  args[6].Float(),
		Width:  args[7].Int(), Height: args[8].Int(),
		IterationLimit: args[9].Int(),
	}

	// set := mbset.MultibrodSet{
	// 	Seed: 0, Power: 2,
	// 	Origin: 0, Scale: 1,
	// 	Width: 50, Height: 100,
	// 	IterationLimit: 100,
	// }

	set.Generate()

	arr := js.Global().Get("Uint8ClampedArray").New(len(set.Image))
	js.CopyBytesToJS(arr, set.Image)

	arr.Set("height", js.ValueOf(set.Height))
	arr.Set("width", js.ValueOf(set.Width))

	return arr
}

func main() {
	js.Global().Set("generateMandelbrot", js.FuncOf(generateMandelbrot))

	<-make(chan struct{})
}
