package goft

import (
	"fmt"
	"reflect"
)

func TestMultiply(t *testing.T) {

	cases := []struct {
		name     string
		v1       vector
		v2       vector
		expected vector
	}{
		{
			"Succes Ints 1",
			vector{1, 2, 3, 4},
			vector{4, 3, 2, 1},
			vector{4, 6, 6, 4},
		},
		{
			"Success Floats 1.2",
			vector{1, 2, 3.3, 4.7},
			vector{4.2, 3.3, 2, 1},
			vector{4.2, 6.6, 6.6, 4.7},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := tt.v1.multiply(tt.v2)
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Unexpected response:\nGot: %v, Wanted: %v", actual, tt.expected)
			}
		})
	}
}

func TestMatrix(t *testing.T) {
	aData := []float32{1, 0, 0, 1}
	a := Matrix{data: aData, dim: Dimension{rows: 2, cols: 2}}

	bData := []float32{1.222, 0.123123, 0.23123, 1.33333}
	b := Matrix{data: bData, dim: Dimension{rows: 2, cols: 2}}

	m, err := a.multiply(&b)
	if err != nil {
		t.Fatal()
	}
}
