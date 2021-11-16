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

// Linspace 初始化等间距行向量
//
// start 最小值
//
// end 最大值
//
// dim 1 表示列向量；2 表示行向量
func Linspace(start float64, end float64, num int, dim int) Matrix {

	array := make([]float64, num)
	step := (end - start) / float64(num)
	for j := 0; j < num; j++ {
		array = append(array, step*float64(j))
	}

	return NewVector(array, dim)
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

// Det 行列式
func Det(A Matrix) float64 {
	if A.Col != A.Row {
		panic("Det operation: matrix must be square.")
	}

	B := A.Copy()

	m := B.Row
	n := B.Col
	for j := 0; j < m; j++ {
		for i := j + 1; i < m; i++ {
			if B.Get(i, j) != 0 {
				c := B.Get(i, j) / B.Get(j, j)
				// k -> Col
				for k := 0; k < n; k++ {
					v := B.Get(i, k) - c*B.Get(j, k)
					B.Set(i, k, v)
				}
			}
		}
	}

	det := 1.0
	for i := 0; i < m; i++ {
		det *= B.Get(i, i)
	}
	return det
}

// Inv 初等变换求逆矩阵
func Inv(A Matrix) (S Matrix) {

	if Det(A) == 0.0 {
		panic("Det(A) is zero, so the matrix cannot be inv")
	}

	if A.Size() == 1 {
		v := 1.0 / A.Get(0, 0)
		return NewVector([]float64{v}, 1)
	}

	shape := Shape{
		Row: A.Row,
		Col: A.Col * 2,
	}

	B := Zeros(shape)
	for i := 0; i < B.Row; i++ {
		for j := 0; j < B.Col; j++ {
			if j < A.Col {
				B.Set(i, j, A.Get(i, j))
			} else if j-A.Col == i {
				B.Set(i, j, 1)
			} else {
				B.Set(i, j, 0)
			}
		}
	}

	m := B.Row
	n := B.Col

	// 上三角矩阵
	// [a b]
	// [0 c]
	for j := 0; j < m; j++ {
		for i := j + 1; i < m; i++ {
			if B.Get(i, j) != 0 {
				c := B.Get(i, j) / B.Get(j, j)
				// k -> Col
				for k := 0; k < n; k++ {
					v := B.Get(i, k) - c*B.Get(j, k)
					B.Set(i, k, v)
				}
			}
		}
	}

	// 下三角矩阵
	// [a 0]
	// [0 c]
	for j := m - 1; j >= 0; j-- {
		for i := j - 1; i >= 0; i-- {
			if B.Get(i, j) != 0 {
				c := B.Get(i, j) / B.Get(j, j)
				// k -> Col
				for k := 0; k < n; k++ {
					v := B.Get(i, k) - c*B.Get(j, k)
					B.Set(i, k, v)
				}
			}
		}
	}

	// 单位矩阵 [E | b]
	for i := 0; i < m; i++ {
		if B.Get(i, i) != 1 {
			for j := m; j < n; j++ {
				v := B.Get(i, j) / B.Get(i, i)
				B.Set(i, j, v)
			}
		}
	}

	// 复制逆矩阵部分
	S = Zeros(A.Shape)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			S.Set(i, j, B.Get(i, j+m))
		}
	}

	return
}

// LU LU分解
func LU(A Matrix) (U Matrix, L Matrix) {
	if A.Col != A.Row {
		panic("matrix A must be square.")
	}
	m := A.Row
	n := A.Col

	U = A.Copy()
	L = Eye(m)

	for j := 0; j < n-1; j++ {
		for i := j + 1; i < n; i++ {
			v := U.Get(i, j) / U.Get(j, j)
			L.Set(i, j, v)
			for k := j; k < n; k++ {
				v = U.Get(i, k) - L.Get(i, j)*U.Get(j, k)
				U.Set(i, k, v)
			}
		}
	}
	return
}

// Cholesky Cholesky分解
func Cholesky(A Matrix) (U Matrix, L Matrix) {
	if A.Col != A.Row {
		panic("matrix A must be square.")
	}
	n := A.Col

	L = Zeros(A.Shape)
	for j := 0; j < n; j++ {
		v := A.Get(j, j)
		for k := 0; k < j; k++ {
			v = v - L.Get(j, k)*L.Get(j, k)
		}
		if v < 0 {
			panic("Cholesky(A) require that diagonal element of A is positive")
		}
		v = math.Sqrt(v)
		L.Set(j, j, v)
		for i := j + 1; i < n; i++ {
			u := A.Get(i, j)
			for k := 0; k < j; k++ {
				u = u - L.Get(i, k)*L.Get(j, k)
			}
			L.Set(i, j, u/v)
		}
	}

	U = L.T()

	return
}
