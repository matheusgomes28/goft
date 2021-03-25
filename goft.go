package goft

import "fmt"

type vector []float32

func (m *Matrix) unsafeAt(row, col int) *float32 {
	return &m.data[row*m.dim.cols+col]
}

// multiply two vectors of the same size together
func (v1 vector) multiply(v2 vector) (vector, error) {

	v1Len := len(v1)
	v2Len := len(v2)

	if v1Len != v2Len {
		return nil, fmt.Errorf("vectors aren't the same size: v1 size = %d, v2 size = %d", v1Len, v2Len)
	}

	v3 := make(vector, v1Len)
	for i := 0; i < v1Len; i++ {
		v3[i] = v1[i] * v2[i]
	}

	return v3, nil
}
