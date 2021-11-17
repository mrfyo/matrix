package matrix

import "math"

func IsVector(A Matrix) bool {
	return A.Row == 1 || A.Col == 1
}

// Norm 范数
//
// 向量: 2-范数; 矩阵: F-范数
func Norm(A Matrix) (d float64) {
	for i := 0; i < A.Row; i++ {
		for j := 0; j < A.Col; j++ {
			v := A.Get(i, j)
			d += v * v
		}
	}
	d = math.Sqrt(d)
	return
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
