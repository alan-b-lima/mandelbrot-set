package mbset

import (
	"math"
	"math/cmplx"

	"github.com/alan-b-lima/mandelbrot-set/src/internal/color"
)

type MultibrodSet struct {
	Seed, Power    complex128
	Origin         complex128
	Scale          float64
	Width, Height  int
	IterationLimit int

	Image []byte
}

const (
	Bound = 2
)

func (s *MultibrodSet) Generate() {
	if s.Image != nil {
		return
	}

	s.Image = make([]byte, 4*s.Height*s.Width)

	for y := range s.Height {
		offset := y * 4 * s.Width
		row := s.Image[offset : offset+4*s.Width]

		for x := range s.Width {
			offset := 4 * x
			cell := row[offset : offset+4]

			c := s.MapPoint(x, y)
			cell[0], cell[1], cell[2] = s.GeneratePoint(c)
			cell[3] = 255
		}
	}
}

func (s *MultibrodSet) GeneratePoint(c complex128) (uint8, uint8, uint8) {
	z := s.Seed
	for i := range s.IterationLimit {
		if real(z)*real(z)+imag(z)*imag(z) >= Bound*Bound {
			return ColorFromLimit(i)
		}

		z = cmplx.Pow(z, s.Power) + c
	}

	return 0, 0, 0
}

func (s *MultibrodSet) MapPoint(x, y int) complex128 {
	scale := math.Pow(2, s.Scale)
	return complex(
		real(s.Origin)+scale*(float64(x-s.Width>>1)),
		imag(s.Origin)+scale*(float64(y-s.Height>>1)),
	)
}

func ColorFromLimit(i int) (R, G, B uint8) {
	return color.HSLToRGB(
		245*(1-min(float64(i), 25)/25),
		1.0,
		0.5,
	)
}
