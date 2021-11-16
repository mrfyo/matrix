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

// Matrix Struct is two-dim matrix like Matlab. 
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

// NewMatrix 默认构造方法
func NewMatrix(shape Shape, array []float64) (A Matrix) {
	A.Shape = shape
	A.array = array
	return
}

// NewSquareMatrix 方块矩阵
func NewSquareMatrix(n int, array []float64) (A Matrix) {
	A.Shape = Shape{n, n}
	A.array = array
	return A
}

// NewVector 新建向量，dim = 1 表示列向量；dim = 2 表示行向量
func NewVector(array []float64, dim int) (A Matrix) {
	n := len(array)
	if dim == 1 {
		A = NewMatrix(Shape{n, 1}, array)
	} else {
		A = NewMatrix(Shape{1, n}, array)
	}

	return
}

