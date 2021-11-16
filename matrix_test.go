package matrix

import "testing"

func TestGet(t *testing.T) {
	A := NewMatrix(Shape{2, 2}, []float64{1, 2, 3, 4})

	if A.Get(0, 0) != 1 {
		t.Error("error access by Get method.")
	}

	if A.Get(0, 1) != 2 {
		t.Error("error access by Get method.")
	}

	if A.Get(1, 0) != 3 {
		t.Error("error access by Get method.")
	}
	if A.Get(1, 1) != 4 {
		t.Error("error access by Get method.")
	}

}

func TestGetIndex(t *testing.T) {
	A := NewVector([]float64{1, 2, 3, 4}, 1)

	if A.GetIndex(2) != 3 {
		t.Error("error access by GetIndex method.")
	}
}

func TestSet(t *testing.T) {
	A := NewVector([]float64{1, 2, 3, 4}, 1)

	A.Set(0, 0, 5)

	if A.Get(0, 0) != 5 {
		t.Error("error method: Set")
	}

}

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
