package main

import (
	"math"
)

type Sphere struct {
	Radius   float64
	Position Vec

	Emission       Color
	Color          Color
	ReflectionType ReflectionType
}

func (s Sphere) Intersect(ray Ray) (isIntersect bool, hitpoint Hitpoint) {
	po := s.Position.SubVec(ray.Org)
	b := po.Dot(ray.Dir)
	d4 := b*b - po.Dot(po) + s.Radius*s.Radius
	if d4 < 0.0 {
		isIntersect = false
		return
	}

	sqrtD4 := math.Sqrt(d4)
	t1, t2 := b-sqrtD4, b+sqrtD4
	if t1 < Eps && t2 < Eps {
		isIntersect = false
		return
	} else if t1 > Eps {
		hitpoint.Distance = t1
	} else {
		hitpoint.Distance = t2
	}

	hitpoint.Position = ray.Org.AddVec(ray.Dir.Mul(hitpoint.Distance))
	hitpoint.Normal = hitpoint.Position.SubVec(s.Position).Normalize()
	hitpoint.Emission = s.Emission
	hitpoint.Color = s.Color
	hitpoint.ReflectionType = s.ReflectionType

	isIntersect = true
	return
}
