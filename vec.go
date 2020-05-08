package main

import "math"

type Vec struct {
	X, Y, Z float64
}

func (a Vec) AddVec(b Vec) Vec {
	return Vec{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func (a Vec) SubVec(b Vec) Vec {
	return Vec{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (a Vec) MulVec(b Vec) Vec {
	return Vec{
		X: a.X * b.X,
		Y: a.Y * b.Y,
		Z: a.Z * b.Z,
	}
}

func (a Vec) DivVec(b Vec) Vec {
	return Vec{
		X: a.X / b.X,
		Y: a.Y / b.Y,
		Z: a.Z / b.Z,
	}
}

func (a Vec) LengthSquared() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

func (a Vec) Length() float64 {
	return math.Sqrt(a.LengthSquared())
}

func (a Vec) Mul(f float64) Vec {
	return Vec{
		X: f * a.X,
		Y: f * a.Y,
		Z: f * a.Z,
	}
}

func (a Vec) Dot(b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec) Normalize() Vec {
	return a.Mul(1.0 / a.Length())
}

func (a Vec) Cross(b Vec) Vec {
	return Vec{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}
