package matrix

import (
	"fmt"
	"math"
	"strings"
)

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

type Matrix struct {
	Shape
	array []float64
}

// Get 获取元素
func (A Matrix) Get(i, j int) float64 {
	if i >= A.Row {
		panic(fmt.Sprintf("index out of bounds: i[%d] >= Row[%d]", i, A.Row))
	}

	if j >= A.Col {
		panic(fmt.Sprintf("index out of bounds: j[%d] >= Cow[%d]", j, A.Col))
	}

	ind := i*A.Col + j
	return A.array[ind]
}

// GetIndex 按下标索取
func (A Matrix) GetIndex(ind int) float64 {
	if ind >= len(A.array) {
		panic(fmt.Sprintf("index out of bounds: %d >= %d", ind, len(A.array)))
	}
	return A.array[ind]
}

func (A Matrix) String() string {
	var Cols []string
	for i := 0; i < A.Row; i++ {
		var Col []string
		for j := 0; j < A.Col; j++ {
			v := A.Get(i, j)
			if v-math.Floor(v) < 1e-6 {
				Col = append(Col, fmt.Sprintf("%d.", int(v)))
			} else {
				Col = append(Col, fmt.Sprintf("%5f", A.Get(i, j)))
			}

		}
		Cols = append(Cols, strings.Join(Col, ", "))
	}
	return fmt.Sprintf("[%s]", strings.Join(Cols, "; "))
}
