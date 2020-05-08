package main

import (
	"math"
	"math/rand"
)

const (
	DepthMin   = 5
	DepthLimit = 64
)

var (
	BackgroundColor = Color{0.0, 0.0, 0.0}
)

func radiance(ray Ray, rnd *rand.Rand) Color {
	return _radiance(ray, rnd, 1)
}

func _radiance(ray Ray, rnd *rand.Rand, depth int) Color {
	ok, hitpoint := intersectScene(ray)
	if !ok {
		return BackgroundColor
	}

	var orientedNormal Vec
	if hitpoint.Normal.Dot(ray.Dir) < 0.0 {
		orientedNormal = hitpoint.Normal
	} else {
		orientedNormal = hitpoint.Normal.Mul(-1.0)
	}

	var russianRouletteProbability float64
	if depth <= DepthMin {
		russianRouletteProbability = 1.0
	} else {
		russianRouletteProbability = math.Max(math.Max(hitpoint.Color.R, hitpoint.Color.G), hitpoint.Color.B)

		if depth > DepthLimit {
			russianRouletteProbability *= math.Pow(0.5, float64(depth-DepthLimit))
		}

		if rnd.Float64() >= russianRouletteProbability {
			return hitpoint.Emission
		}
	}

	var incomingRadiance Color
	var weight Color

	switch hitpoint.ReflectionType {
	case ReflectionTypeDiffuse:
		var w, u, v Vec
		w = orientedNormal
		if math.Abs(w.X) > Eps {
			u = Vec{0.0, 1.0, 0.0}.Cross(w).Normalize()
		} else {
			u = Vec{1.0, 0.0, 0.0}.Cross(w).Normalize()
		}
		v = w.Cross(u)

		r1 := Pi2 * rnd.Float64()
		r2 := rnd.Float64()
		dir := ((u.Mul(math.Cos(r1) * math.Sqrt(r2))).
			AddVec(v.Mul(math.Sin(r1) * math.Sqrt(r2))).
			AddVec(w.Mul(math.Sqrt(1.0 - r2)))).
			Normalize()

		incomingRadiance = _radiance(Ray{hitpoint.Position, dir}, rnd, depth+1)
		weight = hitpoint.Color.Div(russianRouletteProbability)
	case ReflectionTypeSpecular:
		incomingRadiance = _radiance(
			Ray{
				hitpoint.Position,
				ray.Dir.SubVec(hitpoint.Normal.Mul(2.0 * hitpoint.Normal.Dot(ray.Dir))),
			},
			rnd,
			depth+1,
		)
		weight = hitpoint.Color.Div(russianRouletteProbability)
	case ReflectionTypeRefraction:
		reflectionRay := Ray{
			hitpoint.Position,
			ray.Dir.SubVec(hitpoint.Normal.Mul(2.0 * hitpoint.Normal.Dot(ray.Dir))),
		}
		isIntoShape := hitpoint.Normal.Dot(orientedNormal) > 0.0

		// Snell's law
		var (
			nc    float64 = 1.0
			nt    float64 = Ri
			nnt   float64
			ddn   float64 = ray.Dir.Dot(orientedNormal)
			cos2t float64
		)
		if isIntoShape {
			nnt = nc / nt
		} else {
			nnt = nt / nc
		}
		cos2t = 1.0 - nnt*nnt*(1.0-ddn*ddn)

		// total reflection
		if cos2t < 0.0 {
			incomingRadiance = _radiance(reflectionRay, rnd, depth+1)
			weight = hitpoint.Color.Div(russianRouletteProbability)
			break
		}

		refractionRay := Ray{Org: hitpoint.Position}
		if isIntoShape {
			refractionRay.Dir = ray.Dir.Mul(nnt).SubVec(hitpoint.Normal.Mul(ddn*nnt + math.Sqrt(cos2t)))
		} else {
			refractionRay.Dir = ray.Dir.Mul(nnt).AddVec(hitpoint.Normal.Mul(ddn*nnt + math.Sqrt(cos2t)))
		}

		var (
			r0 float64 = ((nt - nc) * (nt - nc)) / ((nt + nc) * (nt + nc))
			c  float64
			re float64
			tr float64
		)
		if isIntoShape {
			c = 1.0 + ddn
		} else {
			c = 1.0 - refractionRay.Dir.Dot(orientedNormal.Mul(-1.0))
		}
		re = r0 + (1.0-r0)*math.Pow(c, 5.0)
		tr = (1.0 - re) * nnt * nnt

		if depth <= 3 {
			ir1 := _radiance(reflectionRay, rnd, depth+1).Mul(re)
			ir2 := _radiance(refractionRay, rnd, depth+1).Mul(tr)
			incomingRadiance = ir1.AddColor(ir2)
			weight = hitpoint.Color.Div(russianRouletteProbability)
			break
		}

		probability := 0.25 + 0.5*re
		if rnd.Float64() < probability {
			incomingRadiance = _radiance(reflectionRay, rnd, depth+1).Mul(re)
			weight = hitpoint.Color.Div(probability * russianRouletteProbability)
		} else {
			incomingRadiance = _radiance(refractionRay, rnd, depth+1).Mul(tr)
			weight = hitpoint.Color.Div((1.0 - probability) * russianRouletteProbability)
		}
	}

	return hitpoint.Emission.AddColor(weight.MulColor(incomingRadiance))
}
