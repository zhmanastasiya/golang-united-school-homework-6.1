package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

func (t Triangle) CalcArea() float64 {
	a := (t.Side * 3) / 2
	return math.Sqrt(a * (a - t.Side) * (a - t.Side) * (a - t.Side))
}
