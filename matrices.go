package goft

import "fmt"

type Dimension struct {
	rows, cols int
}

type Matrix struct {
	data []float32
	dim  Dimension
}

func ZeroMat(rows, cols int) Matrix {
	return Matrix{
		data: make([]float32, rows*cols),
		dim:  Dimension{rows: rows, cols: cols},
	}
}

func (a *Matrix) multiply(b *Matrix) (Matrix, error) {
	if a.dim.cols != b.dim.rows {
		return Matrix{}, fmt.Errorf("Inner dimension mismatch: a is (%d, %d), b is (%d, %d)",
			a.dim.rows,
			a.dim.cols,
			b.dim.rows,
			b.dim.cols)
	}

	result_matrix := Matrix{data: make([]float32, a.dim.rows*b.dim.cols),
		dim: Dimension{a.dim.rows, b.dim.cols}}
	for row := 0; row < a.dim.rows; row++ {
		for col := 0; col < b.dim.cols; col++ {

			destination := result_matrix.unsafeAt(row, col)
			for i := 0; i < a.dim.cols; i++ {
				*destination += (a.data[row*a.dim.cols+i] * b.data[i*b.dim.cols+col])
			}
		}
	}

	return result_matrix, nil
}
