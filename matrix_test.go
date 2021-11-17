package matrix

import "testing"

func TestGet(t *testing.T) {
	A := Builder().Row().Link(1, 2).Link(3, 4).Build()

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
	A := Builder().Row().Link(1, 2, 3, 4).Build()

	if A.GetIndex(2) != 3 {
		t.Error("error access by GetIndex method.")
	}
}

func TestSet(t *testing.T) {
	A := Builder().Row().Link(1, 2).Link(3, 4).Build()

	A.Set(0, 0, 5)

	if A.Get(0, 0) != 5 {
		t.Error("error method: Set")
	}

}
