package features

import (
	"github.com/nishantb06/ray_tracing_challenge_golang/utils"
)

const epsilon = 0.00001

type Matrix struct {
	Rows   int
	Cols   int
	Values [][]float64
}

func Matrix_(rows, cols int) Matrix {
	m := Matrix{Rows: rows, Cols: cols}
	m.Values = make([][]float64, rows)
	for i := 0; i < rows; i++ {
		m.Values[i] = make([]float64, cols)
	}
	return m
}

func (m Matrix) Get(row, col int) float64 {
	return m.Values[row][col]
}

func (m *Matrix) Set(row, col int, value float64) {
	m.Values[row][col] = value
}

func (m Matrix) FillValues(values ...float64) {
	for i, value := range values {
		row := i / m.Cols
		col := i % m.Cols
		m.Set(row, col, value)
	}
}

func (m Matrix) Equals(m2 Matrix) bool {
	if m.Rows != m2.Rows || m.Cols != m2.Cols {
		return false
	}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if !utils.Equal(m.Get(i, j), m2.Get(i, j)) {
				return false
			}
		}
	}
	return true
}
