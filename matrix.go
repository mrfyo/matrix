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

func (s Shape) Size() int {
	return s.Row * s.Col
}

func (s Shape) String() string {
	return fmt.Sprintf("(%d, %d)", s.Row, s.Col)
}

func ShapeNotEqual(a, b Shape) bool {
	return a.Row != b.Row || a.Col != b.Col
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

// Set 设置元素
func (A Matrix) Set(i, j int, v float64) {
	ind := i*A.Col + j
	A.array[ind] = v
}

// SetIndex 按下标替换
func (A Matrix) SetIndex(ind int, v float64) {
	A.array[ind] = v
}

func (A Matrix) String() string {
	var Cols []string
	for i := 0; i < A.Row; i++ {
		var Col []string
		for j := 0; j < A.Col; j++ {
			v := A.Get(i, j)
			if v-math.Floor(v) < 1e-6 {
				Col = append(Col, fmt.Sprintf("%d", int(v)))
			} else {
				Col = append(Col, fmt.Sprintf("%5f", A.Get(i, j)))
			}

		}
		Cols = append(Cols, strings.Join(Col, ", "))
	}
	return fmt.Sprintf("[%s]", strings.Join(Cols, "; "))
}

// GetCol 获取列向量
func (A Matrix) GetCol(j int) (V Matrix) {
	shape := Shape{
		Row: A.Row,
		Col: 1,
	}

	V = NewMatrix(shape, make([]float64, shape.Size()))
	for i := 0; i < A.Row; i++ {
		V.Set(i, 0, A.Get(i, j))
	}

	return
}

// SetCol 指定位置替换列向量
func (A Matrix) SetCol(j int, V Matrix) {
	for i := 0; i < A.Row; i++ {
		A.Set(i, j, V.Get(i, 0))
	}
}

// GetRow 获取行向量
func (A Matrix) GetRow(i int) (V Matrix) {
	shape := Shape{
		Row: 1,
		Col: A.Col,
	}

	V = NewMatrix(shape, make([]float64, shape.Size()))
	for j := 0; j < shape.Col; j++ {
		V.Set(0, j, A.Get(i, j))
	}

	return
}

// SetRow  指定位置替换行向量
func (A Matrix) SetRow(i int, V Matrix) {
	for j := 0; j < A.Col; j++ {
		A.Set(i, j, V.Get(0, j))
	}
}

// Add 矩阵相加
func (A Matrix) Add(B Matrix) (S Matrix) {
	if ShapeNotEqual(A.Shape, B.Shape) {
		panic(fmt.Sprintf("two matrix cannot [add]. %v x %v", A.Shape, B.Shape))
	}

	S = Zeros(A.Shape)
	for i := 0; i < S.Row; i++ {
		for j := 0; j < S.Col; j++ {
			S.Set(i, j, A.Get(i, j)+B.Get(i, j))
		}
	}

	return
}

// Sub 矩阵相减
func (A Matrix) Sub(B Matrix) (S Matrix) {
	if ShapeNotEqual(A.Shape, B.Shape) {
		panic(fmt.Sprintf("two matrix cannot [sub]. %v x %v", A.Shape, B.Shape))
	}

	S = Zeros(A.Shape)
	for i := 0; i < S.Row; i++ {
		for j := 0; j < S.Col; j++ {
			S.Set(i, j, A.Get(i, j)-B.Get(i, j))
		}
	}

	return
}

// Mul 点乘(同位置相乘，形状不变)
func (A Matrix) Mul(B Matrix) (S Matrix) {

	if ShapeNotEqual(A.Shape, B.Shape) {
		panic(fmt.Sprintf("two matrix cannot [mul]. %v x %v", A.Shape, B.Shape))
	}

	S = Zeros(A.Shape)
	for i := 0; i < S.Row; i++ {
		for j := 0; j < S.Col; j++ {
			S.Set(i, j, A.Get(i, j)*B.Get(i, j))
		}
	}
	return
}

// Dot 矩阵乘法
func (A Matrix) Dot(B Matrix) (S Matrix) {
	if A.Col != B.Row {
		panic(fmt.Sprintf("two matrix cannot [dot]. %v x %v", A.Shape, B.Shape))
	}

	shape := Shape{
		Row: A.Row,
		Col: B.Col,
	}

	S = Zeros(shape)
	for i := 0; i < shape.Row; i++ {
		for j := 0; j < shape.Col; j++ {
			v := 0.0
			for k := 0; k < A.Col; k++ {
				v = v + A.Get(i, k)*B.Get(k, j)
			}
			S.Set(i, j, v)
		}
	}
	return
}

// ScaleMul 矩阵比例乘
func (A Matrix) ScaleMul(k float64) (S Matrix) {

	S = Zeros(A.Shape)

	for i := 0; i < S.Row; i++ {
		for j := 0; j < S.Col; j++ {
			S.Set(i, j, A.Get(i, j)*k)
		}
	}

	return
}

// T 转置
func (A Matrix) T() (S Matrix) {
	shape := Shape{
		Row: A.Col,
		Col: A.Row,
	}
	S = Zeros(shape)

	for i := 0; i < shape.Row; i++ {
		for j := 0; j < shape.Col; j++ {
			S.Set(i, j, A.Get(j, i))
		}
	}
	return
}

// Copy 矩阵复制
func (A Matrix) Copy() (S Matrix) {
	S = Zeros(A.Shape)

	for i := 0; i < S.Row; i++ {
		for j := 0; j < S.Col; j++ {
			S.Set(i, j, A.Get(i, j))
		}
	}
	return
}

func MatrixAdd(target, sour Matrix) {
	shape := target.Shape

	for i := 0; i < shape.Row; i++ {
		for j := 0; j < shape.Col; j++ {
			v := target.Get(i, j) + sour.Get(i, j)
			target.Set(i, j, v)
		}
	}
}

func MatrixSub(target, sour Matrix) {
	shape := target.Shape

	for i := 0; i < shape.Row; i++ {
		for j := 0; j < shape.Col; j++ {
			v := target.Get(i, j) - sour.Get(i, j)
			target.Set(i, j, v)
		}
	}
}

func MatrixScaleMul(A Matrix, k float64) {
	for i := 0; i < A.Row; i++ {
		for j := 0; j < A.Col; j++ {
			v := A.Get(i, j) * k
			A.Set(i, j, v)
		}
	}
}

func MatrixEqual(A, B Matrix) bool {
	if ShapeNotEqual(A.Shape, B.Shape) {
		return false
	}

	for i := 0; i < A.Row; i++ {
		for j := 0; j < A.Col; j++ {
			if math.Abs(A.Get(i, j)-B.Get(i, j)) > 1e-8 {
				return false
			}
		}
	}

	return true
}

// Zeros 零矩阵
func Zeros(shape Shape) Matrix {
	array := make([]float64, shape.Size())
	return NewMatrix(shape, array)
}

// Ones 全 1 矩阵
func Ones(shape Shape) Matrix {
	array := make([]float64, shape.Size())
	for i := 0; i < len(array); i++ {
		array[i] = 1
	}
	return NewMatrix(shape, array)
}

// Full 填充矩阵
func Full(shape Shape, v float64) Matrix {
	array := make([]float64, shape.Size())
	for i := 0; i < len(array); i++ {
		array[i] = v
	}
	return NewMatrix(shape, array)
}

// Eye 单位矩阵
func Eye(n int) Matrix {
	shape := Shape{n, n}
	array := make([]float64, shape.Size())

	m := NewMatrix(shape, array)
	for i := 0; i < n; i++ {
		m.Set(i, i, 1)
	}
	return m
}

// Diag 分块矩阵
func Diag(vs []float64) Matrix {
	n := len(vs)
	A := Zeros(Shape{n, n})
	for i := 0; i < n; i++ {
		A.Set(i, i, vs[i])
	}
	return A
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
