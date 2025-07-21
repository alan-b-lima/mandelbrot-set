package color

import "testing"

type (
	TestData struct {
		given    HSL
		expected RGB
	}

	HSL struct{ H, S, L float64 }
	RGB struct{ R, G, B uint8 }
)

func TestHSLToRGB(t *testing.T) {
	data := []TestData{
		{HSL{0, .0, .0}, RGB{0, 0, 0}},
		{HSL{0, .0, 1}, RGB{255, 255, 255}},
		{HSL{0, 1, .5}, RGB{255, 0, 0}},
		{HSL{120, 1, .5}, RGB{0, 255, 0}},
		{HSL{240, 1, .5}, RGB{0, 0, 255}},
		{HSL{60, 1, .5}, RGB{255, 255, 0}},
		{HSL{180, 1, .5}, RGB{0, 255, 255}},
		{HSL{300, 1, .5}, RGB{255, 0, 255}},
		{HSL{0, .0, .75}, RGB{191, 191, 191}},
		{HSL{0, .0, .5}, RGB{127, 127, 127}},
		{HSL{0, 1, .25}, RGB{127, 0, 0}},
		{HSL{60, 1, .25}, RGB{127, 127, 0}},
		{HSL{120, 1, .25}, RGB{0, 127, 0}},
		{HSL{300, 1, .25}, RGB{127, 0, 127}},
		{HSL{180, 1, .25}, RGB{0, 127, 127}},
		{HSL{240, 1, .25}, RGB{0, 0, 127}},
	}

	for i, test := range data {
		R, G, B := HSLToRGB(test.given.H, test.given.S, test.given.L)
		if test.expected.R != R || test.expected.G != G || test.expected.B != B {
			t.Errorf("Test %d failed: expected %v, got {%d, %d, %d} from %v",
				i, test.expected, R, G, B, test.given)
		}
	}
}
