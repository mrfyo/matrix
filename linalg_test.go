package matrix

import (
	"math"
	"testing"
)

func TestDet(t *testing.T) {
	A := Builder().Row().Link(1, 0).Link(0, 4).Build()

	if Det(A) != 4 {
		t.Error("error method: Det.")
	}

	A = Builder().Row().Link(1, 2).Link(3, 4).Build()

	if Det(A) != -2 {
		t.Error("error method: Det.")
	}
}

func TestInv(t *testing.T) {
	A := Builder().Row().Link(1, 0).Link(0, 4).Build()
	B := Inv(A)

	if B.Get(1, 1) != 0.25 {
		t.Error("error method: Det.")
	}

	A = Builder().Row().Link(1, 2).Link(3, 4).Build()
	B = Inv(A)
	C := []float64{-2, 1, 1.5, -0.5}

	for i := 0; i < B.Size(); i++ {
		if B.GetIndex(i) != C[i] {
			t.Error("error method: Det.")
		}
	}

}

func TestQR(t *testing.T) {
	A := Builder().Row().Link(1, 2, 4).Link(0, 0, 5).Link(0, 3, 6).Build()
	QExpected := Builder().Row().Link(1, 0, 0).Link(0, 0, 1).Link(0, 1, 0).Build()
	RExpected := Builder().Row().Link(1, 2, 4).Link(0, 3, 6).Link(0, 0, 5).Build()
	Q, R := QR(A)

	if !MatrixEqual(QExpected, Q) || !MatrixEqual(RExpected, R) {
		t.Error("error method: QR")
	}

}

func TestQR2(t *testing.T) {
	A := NewSquareMatrix(3, []float64{0, 1, 1, 1, 1, 0, 1, 0, 1})
	QExpected := Builder().Row().Link(0, math.Sqrt(4.0/6), math.Sqrt(1.0/3)).Link(math.Sqrt(1.0/2), math.Sqrt(1.0/6), -math.Sqrt(1.0/3)).Link(math.Sqrt(1.0/2), -math.Sqrt(1.0/6), math.Sqrt(1.0/3)).Build()
	RExpected := Builder().Row().Link(math.Sqrt(2), math.Sqrt(1.0/2), math.Sqrt(1.0/2)).Link(0, math.Sqrt(6.0/4), math.Sqrt(6.0)/6).Link(0, 0, math.Sqrt(4.0/3)).Build()
	Q, R := QR(A)

	if !MatrixEqual(QExpected, Q) || !MatrixEqual(RExpected, R) {
		t.Error("error method: QR")
	}
}
