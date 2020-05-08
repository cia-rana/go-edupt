package main

const Ri = 1.5 // refractive index

type ReflectionType int

const (
	_ ReflectionType = iota
	ReflectionTypeDiffuse
	ReflectionTypeSpecular
	ReflectionTypeRefraction
)
