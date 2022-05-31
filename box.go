package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	errorCapacity = errors.New("out of the shapesCapacity")
	errorIndex    = errors.New("the index is out of range")
	errorCircls   = errors.New("circles are not exist in the list")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("%w", errorCapacity)
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i <= len(b.shapes)-1 {
		for l := range b.shapes {
			if l == i {
				return b.shapes[l], nil
			}
		}
	}
	return nil, fmt.Errorf("%w", errorIndex)
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i <= len(b.shapes)-1 {
		s := b.shapes[i]
		b.shapes = append(b.shapes[i:], b.shapes[i+1:]...)
		return s, nil
	}
	return nil, fmt.Errorf("%w", errorIndex)
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i <= len(b.shapes) {
		s := b.shapes[i]
		b.shapes[i] = shape
		return s, nil
	}
	return nil, fmt.Errorf("%w", errorIndex)
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for i := 0; i < len(b.shapes); i++ {
		sum += Shape.CalcPerimeter(b.shapes[i])
	}
	if len(b.shapes) == 0 {
		return 0
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var area float64
	for i := 0; i < len(b.shapes); i++ {
		area += b.shapes[i].CalcArea()
		area += Shape.CalcArea(b.shapes[i])
	}
	if len(b.shapes) == 0 {
		return 0
	}
	return area
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	newShapes := make([]Shape, 0)
	for i := range b.shapes {
		_, ok := b.shapes[i].(*Circle)
		if !ok {
			newShapes = append(newShapes, b.shapes[i])
		}
	}
	if len(newShapes) == len(b.shapes) {
		return fmt.Errorf("%w", errorCircls)
	}
	b.shapes = newShapes
	return nil
}
