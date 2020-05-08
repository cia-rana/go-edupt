package main

type Shape interface{
	Intersect(Ray) (bool, Hitpoint)
}
