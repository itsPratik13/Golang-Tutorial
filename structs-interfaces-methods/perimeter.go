package main

import "math"

func Perimeter(r Rectangle) float64 {
	perimeter := 2 * (r.l + r.b)
	return perimeter

}
func Area(r Rectangle) float64 {
	return r.l * r.b
}
func AreaCircle(c Circle) float64 {
	return math.Pi*c.r*c.r
}
