package main

import (
	"math"
)

const InvG = 1 / 2.2 // inverse of gamma value

type Color struct {
	R, G, B float64
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = c.adjustGamma(c.R)
	g = c.adjustGamma(c.G)
	b = c.adjustGamma(c.B)
	a = 0xffff
	return
}

func (c Color) adjustGamma(x float64) uint32 {
	return uint32(math.RoundToEven(math.Pow(c.clamp(x), InvG) * 0xffff))
}

func (Color) clamp(x float64) float64 {
	if x < 0.0 {
		return 0.0
	} else if x > 1.0 {
		return 1.0
	}
	return x
}

func (a Color) AddColor(b Color) Color {
	return Color{
		R: a.R + b.R,
		G: a.G + b.G,
		B: a.B + b.B,
	}
}

func (a Color) SubColor(b Color) Color {
	return Color{
		R: a.R - b.R,
		G: a.G - b.G,
		B: a.B - b.B,
	}
}

func (a Color) MulColor(b Color) Color {
	return Color{
		R: a.R * b.R,
		G: a.G * b.G,
		B: a.B * b.B,
	}
}

func (a Color) DivColor(b Color) Color {
	return Color{
		R: a.R / b.R,
		G: a.G / b.G,
		B: a.B / b.B,
	}
}

func (a Color) Mul(f float64) Color {
	return Color{
		R: a.R * f,
		G: a.G * f,
		B: a.B * f,
	}
}

func (a Color) Div(f float64) Color {
	return Color{
		R: a.R / f,
		G: a.G / f,
		B: a.B / f,
	}
}
