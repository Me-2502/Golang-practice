package myutil

import "fmt"

type shape2D interface {
	perimiter() int
	area() int
}

type shape3D interface {
	shape2D
	volume() int
}

type rectangle struct {
	length  int
	breadth int
}

func (r *rectangle) perimiter() int {
	return (r.length + r.breadth) * 2
}

func (r *rectangle) area() int {
	return r.length * r.breadth
}

type square struct {
	length int
}

func (s *square) perimiter() int {
	return s.length * 2
}

func (s *square) area() int {
	return s.length * s.length
}

type cuboid struct {
	length  int
	breadth int
	width   int
}

func (c *cuboid) perimiter() int {
	return (c.breadth + c.length + c.width) * 4
}

func (c *cuboid) area() int {
	return ((c.breadth * c.length) + (c.length * c.width) + (c.width * c.breadth)) * 2
}

func (c *cuboid) volume() int {
	return c.length * c.breadth * c.width
}

func print2D(s shape2D) {
	fmt.Printf("The given shape is a 2D with perimeter %d and area %d\n", s.perimiter(), s.area())
}

func print3D(s shape3D) {
	fmt.Printf("The given shape is a 3D with perimeter %d, area %d and volume %d\n", s.perimiter(), s.area(), s.volume())
}

func Interfaces() {
	// INTERFACES
	var s2d shape2D
	s2d = &rectangle{5, 6}
	print2D(s2d)
	s2d = &square{5}
	print2D(s2d)
	s3d := &cuboid{4, 5, 6}
	print3D(s3d)
	fmt.Println()
}
