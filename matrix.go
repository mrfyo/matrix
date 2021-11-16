package matrix

import "fmt"

type Shape struct {
	Row int
	Col int
}

func ShapeNotEqual(a, b Shape) bool {
	return a.Row != b.Row || a.Col != b.Col
}

func (s Shape) Size() int {
	return s.Row * s.Col
}

func (s Shape) String() string {
	return fmt.Sprintf("(%d, %d)", s.Row, s.Col)
}