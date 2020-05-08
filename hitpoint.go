package main

type Hitpoint struct {
	Distance float64
	Normal   Vec
	Position Vec

	Emission       Color
	Color          Color
	ReflectionType ReflectionType
}
