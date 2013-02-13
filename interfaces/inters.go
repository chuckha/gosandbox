package main

import (
	"fmt"
	"math"
	"reflect"
)

type Polyhedroner interface {
	Volume() float64
	SurfaceArea() float64
}

type Sphere struct {
	radius float64
}

func (s Sphere) Volume() float64 {
	return math.Pi * math.Pow(s.radius, 3.0) * 4.0 / 3.0
}

func (s Sphere) SurfaceArea() float64 {
	return math.Pi * 4.0 * math.Pow(s.radius, 2.0)
}

type Cube struct {
	side float64
}

func (c Cube) Volume() float64 {
	return math.Pow(c.side, 3.0)
}

func (c Cube) SurfaceArea() float64 {
	return 8 * math.Pow(c.side, 2.0)
}

func main() {
	s := Sphere{2.5}
	c := Cube{2.5}
	polyhedrons := []Polyhedroner{s, c}
	for _, v := range polyhedrons {
		fmt.Printf("%v:\n", reflect.TypeOf(v))
		fmt.Printf("Volume: %v\n", v.Volume())
		fmt.Printf("Surface Area: %v\n", v.SurfaceArea())
	}
}
