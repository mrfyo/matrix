package matrix

type matrixBuilder struct {
	row   int
	col   int
	array []float64
}

type matrixRowBuilder struct {
	builder matrixBuilder
}

type matrixColBuilder struct {
	builder matrixBuilder
}

func (b matrixBuilder) Row() matrixRowBuilder {
	return matrixRowBuilder{b}
}

func (b matrixBuilder) Col() matrixColBuilder {
	return matrixColBuilder{b}
}

func (b matrixRowBuilder) Link(v ...float64) matrixRowBuilder {
	if len(v) == 0 {
		panic("len(v) must > 0")
	}

	col := b.builder.col
	if col == 0 {
		col = len(v)
	} else if col < len(v) {
		panic("the number of linked row and first row must euqal")
	}

	b.builder.row++
	b.builder.col = col
	b.builder.array = append(b.builder.array, v...)

	return b
}

func (b matrixColBuilder) Link(v ...float64) matrixColBuilder {
	if len(v) == 0 {
		panic("len(v) must > 0")
	}

	row := b.builder.row
	if row == 0 {
		row = len(v)
	} else if row != len(v) {
		panic("the number of linked col and first col must euqal")
	}
	b.builder.col++
	b.builder.row = row
	b.builder.array = append(b.builder.array, v...)

	return b
}

func (b matrixRowBuilder) Build() Matrix {
	array := b.builder.array
	row := b.builder.row
	col := b.builder.col
	shape := Shape{Row: row, Col: col}
	return NewMatrix(shape, array)
}

func (b matrixColBuilder) Build() Matrix {
	array := b.builder.array
	col := b.builder.col
	row := b.builder.row

	m := Zeros(Shape{Row: row, Col: col})

	for j := 0; j < col; j++ {
		for i := 0; i < row; i++ {
			ind := j*row + i
			m.Set(i, j, array[ind])
		}
	}

	return m
}

func Builder() matrixBuilder {
	return matrixBuilder{}
}
