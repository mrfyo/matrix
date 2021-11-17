package matrix

import "math"

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
