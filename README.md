# matrix
simple matrix in go. lib supported basic matrix operations, linag operations(like `Det`, `Inv`), matrix decomposition(like `LU`, `QR`, `Cholesky`), vector operations(like `Norm`, `Inner`, `Cross`), polynomial and so on.

## Installation

1. install matrix
```shell
go get -u github.com/mrfyo/matrix
```

2. import it in your code
```go
import "github.com/mrfyo/matrix"
```

## Quick Start

### Initialize Matrix

```go
package main

import (
	"fmt"

	mat "github.com/mrfyo/matrix"
)

func main() {

	ZM := mat.Zeros(mat.Shape{Row: 2, Col: 2})
	// [0, 0; 0, 0]
	fmt.Println(ZM)

	OM := mat.Ones(mat.Shape{Row: 2, Col: 2})
	// [1, 1; 1, 1]
	fmt.Println(OM)

	R := mat.Builder().Row().
		Link(1, 2, 3).
		Link(4, 5, 6).
		Link(7, 8, 9).Build()

	// [1, 2, 3;
	//	4, 5, 6;
	//	7, 8, 9]
	fmt.Println(R)

	C := mat.Builder().Col().Link(1, 2, 3).Link(4, 5, 6).Link(7, 8, 9).Build()
	// [1, 4, 7;
	//	2, 5, 8;
	//	3, 6, 9]
	fmt.Println(C)

	// col-vector
	V := mat.NewVector([]float64{1, 2, 3, 4}, 1)
	// [1; 2; 3; 4]
	fmt.Println(V)

	// row-vector
	B := mat.NewVector([]float64{1, 2, 3, 4}, 2)
	// [1, 2, 3, 4]
	fmt.Println(B)
}
```

### Matrix Operation

```go
func main() {
	A := mat.Builder().Row().
		Link(1, 0, 0).
		Link(0, 2, 0).
		Link(0, 0, 3).Build()

	B := mat.Builder().Row().
		Link(1, 2, 3).
		Link(4, 5, 6).
		Link(7, 8, 9).Build()

	// [2, 2, 3; 4, 7, 6; 7, 8, 12]
	fmt.Println(A.Add(B))

	// [0, -2, -3; -4, -3, -6; -7, -8, -6]
	fmt.Println(A.Sub(B))

	// [1, 2, 3; 8, 10, 12; 21, 24, 2]
	fmt.Println(A.Dot(B))

	// [1, 4, 7; 2, 5, 8; 3, 6, 9]
	fmt.Println(B.T())

	// [3, 0, 0; 0, 6, 0; 0, 0, 9]
	fmt.Println(A.ScaleMul(3))
}
```

### Linag Operations

linag operations include `Det`, `Inv`

```go

func main() {
	A := mat.Builder().Row().
		Link(1, 0, 0).
		Link(0, 2, 0).
		Link(0, 0, 3).Build()

	// 6
	fmt.Println(mat.Det(A))

	// [1, 0, 0; 0, 0.500000, 0; 0, 0, 0.333333]
	fmt.Println(mat.Inv(A))
	

}

```

### Matrix Decomposition
Matrix Decomposition include `LU`, `QR`, `Cholesky`


```go
func main() {
	A := mat.Builder().Row().
		Link(1, 2, 4).
		Link(0, 0, 5).
		Link(0, 3, 6).Build()

	Q, R := mat.QR(A)

	// [1, 0, 0; 0, 0, 1; 0, 1, 0]
	fmt.Println(Q)
	// [1, 2, 4; 0, 3, 6; 0, 0, 5]
	fmt.Println(R)

	B := mat.Builder().Row().Link(1, 2, 3).Link(2, 5, 7).Link(3, 5, 3).Build()
	L, U := mat.LU(B)

	// [1, 2, 3; 0, 1, 1; 0, 0, -5]
	fmt.Println(L)

	// [1, 0, 0; 2, 1, 0; 3, -1, 1]
	fmt.Println(U)

	C := mat.Builder().Row().Link(4, 12, -16).Link(12, 37, -43).Link(-16, -43, 98).Build()
	L, LT := mat.Cholesky(C)

	// [2, 0, 0; 6, 1, 0; -8, 5, 3]
	// [2, 6, -8; 0, 1, 5; 0, 0, 3]
	fmt.Println(L)
	fmt.Println(LT)
}
```

### Vector Operations

vector operations include `Norm`, `Inner`, `Cross`.

```go
func main() {
	a := mat.NewVector([]float64{1, 2, 2}, 1)
	b := mat.NewVector([]float64{0, 1, 1}, 1)
	
	// 3
	fmt.Println(mat.Norm(a))
	
	// 4
	fmt.Println(mat.Inner(a, b))
	
	// [0; -1; 1]
	fmt.Println(mat.Cross(a, b))
}

```