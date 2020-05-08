package main

var shapes = []Shape{
	Sphere{Radius: 1e5, Position: Vec{1e5 + 1.0, 40.8, 81.6}, Emission: Color{}, Color: Color{0.75, 0.25, 0.25}, ReflectionType: ReflectionTypeDiffuse},
	Sphere{Radius: 1e5, Position: Vec{-1e5 + 99.0, 40.8, 81.6}, Emission: Color{}, Color: Color{0.25, 0.25, 0.75}, ReflectionType: ReflectionTypeDiffuse},
	Sphere{Radius: 1e5, Position: Vec{50.0, 40.8, 1e5}, Emission: Color{}, Color: Color{0.75, 0.75, 0.75}, ReflectionType: ReflectionTypeDiffuse},
	Sphere{Radius: 1e5, Position: Vec{50.0, 40.8, -1e5 + 250.0}, Emission: Color{}, Color: Color{}, ReflectionType: ReflectionTypeDiffuse},
	Sphere{Radius: 1e5, Position: Vec{50.0, 1e5, 81.6}, Emission: Color{}, Color: Color{0.75, 0.75, 0.75}, ReflectionType: ReflectionTypeDiffuse},
	Sphere{Radius: 1e5, Position: Vec{50.0, -1e5 + 81.6, 81.6}, Emission: Color{}, Color: Color{0.75, 0.75, 0.75}, ReflectionType: ReflectionTypeDiffuse},
	Sphere{Radius: 20.0, Position: Vec{65.0, 20.0, 20.0}, Emission: Color{}, Color: Color{0.25, 0.75, 0.25}, ReflectionType: ReflectionTypeDiffuse},
	Sphere{Radius: 16.5, Position: Vec{27.0, 16.5, 47.0}, Emission: Color{}, Color: Color{0.99, 0.99, 0.99}, ReflectionType: ReflectionTypeSpecular},
	Sphere{Radius: 16.5, Position: Vec{77.0, 16.5, 78.0}, Emission: Color{}, Color: Color{0.99, 0.99, 0.99}, ReflectionType: ReflectionTypeRefraction},
	Sphere{Radius: 15.0, Position: Vec{50.0, 90.0, 81.6}, Emission: Color{36.0, 36.0, 36.0}, Color: Color{}, ReflectionType: ReflectionTypeDiffuse},
}

func intersectScene(ray Ray) (bool, Hitpoint) {
	hitpoint := Hitpoint{Distance: Inf}
	for _, shape := range shapes {
		if ok, h := shape.Intersect(ray); ok && h.Distance < hitpoint.Distance {
			hitpoint = h
		}
	}

	return hitpoint.Distance < Inf, hitpoint
}
