package matrix

import (
	"math"
	"testing"
)

func TestDet(t *testing.T) {
	A := NewSquareMatrix(2, []float64{1, 0, 0, 4})

	if Det(A) != 4 {
		t.Error("error method: Det.")
	}

	A = NewSquareMatrix(2, []float64{1, 2, 3, 4})

	if Det(A) != -2 {
		t.Error("error method: Det.")
	}
}

func TestInv(t *testing.T) {
	A := NewSquareMatrix(2, []float64{1, 0, 0, 4})
	B := Inv(A)

	if B.Get(1, 1) != 0.25 {
		t.Error("error method: Det.")
	}

	A = NewSquareMatrix(2, []float64{1, 2, 3, 4})
	B = Inv(A)
	C := []float64{-2, 1, 1.5, -0.5}

	for i := 0; i < B.Size(); i++ {
		if B.GetIndex(i) != C[i] {
			t.Error("error method: Det.")
		}
	}

}

func TestQR(t *testing.T) {
	A := NewSquareMatrix(3, []float64{1, 2, 4, 0, 0, 5, 0, 3, 6})
	QExpected := NewSquareMatrix(3, []float64{1, 0, 0, 0, 0, 1, 0, 1, 0})
	RExpected := NewSquareMatrix(3, []float64{1, 2, 4, 0, 3, 6, 0, 0, 5})
	Q, R := QR(A)

	if !MatrixEqual(QExpected, Q) || !MatrixEqual(RExpected, R) {
		t.Error("error method: QR")
	}

}

func TestQR2(t *testing.T) {
	A := NewSquareMatrix(3, []float64{0, 1, 1, 1, 1, 0, 1, 0, 1})
	QExpected := NewSquareMatrix(3,[]float64{0, math.Sqrt(4.0 / 6), math.Sqrt(1.0 / 3),
			math.Sqrt(1.0 / 2), math.Sqrt(1.0 / 6), -math.Sqrt(1.0 / 3),
			math.Sqrt(1.0 / 2), -math.Sqrt(1.0 / 6), math.Sqrt(1.0 / 3)})
	RExpected := NewSquareMatrix(3, []float64{math.Sqrt(2), math.Sqrt(1.0 / 2), math.Sqrt(1.0 / 2),
		0, math.Sqrt(6.0 / 4), math.Sqrt(6.0)/6,
		0, 0, math.Sqrt(4.0 / 3)})
	Q, R := QR(A)

	if !MatrixEqual(QExpected, Q) || !MatrixEqual(RExpected, R) {
		t.Error("error method: QR")
	}
}
