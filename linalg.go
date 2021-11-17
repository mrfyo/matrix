package matrix

import "math"

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

