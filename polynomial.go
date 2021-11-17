package matrix

import "math"

const (
	eps         = 1e-12
	maxValue    = 1e3
	maxTryCount = 10_000
)

// Conv 多项式乘法（离散卷积)
func Conv(F, G Matrix) (Y Matrix) {
	l := F.Size() + G.Size() - 1
	if F.Row == 1 && G.Row == 1 {
		Y = Zeros(Shape{1, l})
	} else if F.Col == 1 && G.Col == 1 {
		Y = Zeros(Shape{l, 1})
	} else {
		panic("conv only support vector")
	}

	for i := 0; i < l; i++ {
		v := 0.0
		for j := 0; j < G.Size(); j++ {
			if i >= j && (i-j) < F.Size() {
				v += F.GetIndex(i-j) * G.GetIndex(j)
			}

		}
		Y.SetIndex(i, v)
	}
	return
}

// Diff 多项式差分
func Diff(A Matrix) (B Matrix) {
	size := A.Size()
	B = Zeros(Shape{1, size - 1})

	for i := 0; i < B.Size(); i++ {
		v := A.GetIndex(i) * float64(size-i+1)
		B.SetIndex(i, v)
	}
	return
}

// PolyEvaluate 多项式求值
// f(x) = x^2 + 2x + 1; f(2) = 4 + 4 + 1 = 9
func PolyEvaluate(F Matrix, x float64) float64 {
	sum := 0.0
	size := F.Size()
	for i := 0; i < size; i++ {
		k := size - i - 1
		sum += math.Pow(x, float64(k)) * F.GetIndex(i)
	}

	return sum
}

// Root 高阶方程近似求解。
// 二阶采用韦达定理，忽略复数根的虚部；
// 三阶采用盛金公式，忽略纯虚数根；
// 四阶采用费拉里法，忽略复数根的虚部；
// 四阶以上采用近似逼近法
func Root(X Matrix) (Y Matrix) {

	var T Matrix
	order := X.Size() - 1
	for i := 0; i < X.Size(); i++ {
		if X.GetIndex(i) == 0 {
			order--
		} else {
			T = Zeros(Shape{1, order + 1})
			for j := 0; j < T.Size(); j++ {
				T.SetIndex(j, X.GetIndex(i+j))
			}
			break
		}
	}

	rs := make([]float64, 0, order)

PM:
	for k := order; k >= 1; k-- {
		switch k {
		case 1:
			a := T.GetIndex(0)
			b := T.GetIndex(1)
			if a*b < 0 {
				p := -b / a
				rs = append(rs, p)
			}
			break PM
		case 2:
			// 求根公式
			a := T.GetIndex(0)
			b := T.GetIndex(1)
			c := T.GetIndex(2)

			delta := b*b - 4*a*c
			if delta == 0 {
				x := (-1 * b) / (2 * a)
				rs = append(rs, x, x)
			} else if delta > 0 {
				x1 := (-1*b - math.Sqrt(delta)) / (2 * a)
				x2 := (-1*b + math.Sqrt(delta)) / (2 * a)
				rs = append(rs, x1, x2)
			} else {
				// 仅考虑复数根的实部
				x := -b / (2 * a)
				rs = append(rs, x, x)
			}
			break PM
		case 3:
			// 盛金公式
			a := T.GetIndex(0)
			b := T.GetIndex(1)
			c := T.GetIndex(2)
			d := T.GetIndex(3)

			A := b*b - 3*a*c
			B := b*c - 9*a*d
			C := c*c - 3*b*d

			delta := B*B - 4*A*C

			if A == B {
				x := -1.0 * c / b
				rs = append(rs, x, x, x)

			} else {
				if delta > 0 {
					// 忽略纯虚根
					y1 := math.Cbrt(A*b + 1.5*a*(-B+math.Sqrt(delta)))
					y2 := math.Cbrt(A*b + 1.5*a*(-B-math.Sqrt(delta)))
					x1 := (-b - (y1 + y2)) / (3 * a)
					// x2 := (-2*b + (y1 + y2) + math.Cbrt(3)*(y1-y2)) / (6 * a)
					// x3 := (-2*b + (y1 + y2) - math.Cbrt(3)*(y1-y2)) / (6 * a)
					rs = append(rs, x1)
				} else if delta == 0 {
					t := B / A
					x1 := -b*a + t
					x2 := -t / 2
					rs = append(rs, x1, x2, x2)
				} else {
					sqrtA := math.Sqrt(A)
					t := (2*A*b - 3*a*B) / (2 * A * sqrtA)
					theta := math.Acos(t)
					cosT := math.Cos(theta / 3.0)
					sinT := math.Sin(theta / 3.0)

					x1 := (-b - 2*sqrtA*cosT) / (3 * a)
					x2 := (-b + sqrtA*(cosT+math.Sqrt(3)*sinT)) / (3 * a)
					x3 := (-b + sqrtA*(cosT-math.Sqrt(3)*sinT)) / (3 * a)

					rs = append(rs, x1, x2, x3)
				}

			}
			break PM
		case 4:
			// 费拉里法
			a := T.GetIndex(0)
			b := T.GetIndex(1) / a
			c := T.GetIndex(2) / a
			d := T.GetIndex(3) / a
			e := T.GetIndex(4) / a

			H := Zeros(Shape{1, 4})
			H.SetIndex(0, 8.0)
			H.SetIndex(1, -4*c)
			H.SetIndex(2, 2*b*d-8*e)
			H.SetIndex(3, -e*(b*b-4*c)-d*d)

			G := Root(H)
			var y, M, N float64
			for j := 0; j < G.Size(); j++ {
				y = G.GetIndex(j)
				M = math.Sqrt(8*y + b*b - 4*c)
				N = b*y - d
				if M > 0 {
					break
				}
			}

			S1 := NewVector([]float64{2, b + M, 2 * (y + N/M)}, 2)
			S2 := NewVector([]float64{2, b - M, 2 * (y - N/M)}, 2)

			r1 := Root(S1)
			r2 := Root(S2)

			rs = append(rs, r1.GetIndex(0), r1.GetIndex(1), r2.GetIndex(0), r2.GetIndex(1))

			break PM
		default:
			low, high := -maxValue, maxValue
			p := 0.0
			i := 0
			for ; i < maxTryCount; i++ {
				sum := PolyEvaluate(T, p)
				if math.Abs(sum) < eps || (math.Abs(low-high) < eps) {
					break
				}

				if sum > 0 {
					if PolyEvaluate(T, low) < 0 {
						high = p
						p = (p + low) / 2.0
					} else {
						if PolyEvaluate(T, high) < 0 {
							low = p
							p = (p + high) / 2.0
						} else {
							break
						}
					}
				} else {
					if PolyEvaluate(T, low) > 0 {
						high = p
						p = (p + low) / 2.0
					} else {
						if PolyEvaluate(T, high) > 0 {
							low = p
							p = (p + high) / 2.0
						} else {
							break
						}
					}
				}
			}

			rs = append(rs, p)
			Q := Zeros(Shape{1, k})
			M := NewVector([]float64{1, -p}, 2)
			for j := 0; j < Q.Size(); j++ {
				a := T.GetIndex(j)
				Q.SetIndex(j, a)
				MatrixSub(T, Conv(Q, M))
			}
			T = Q
		}

	}

	Y = NewVector(rs, 2)
	return
}


