package randAlgs

import "log"

// IndMatrix represents a graph as incidence matrix.
type IndMatrix struct {
	Dim    int
	Values []int // dim * dim values expected
}

// CreateEmptyIndMatrix returns empty IndMatrix of given dimension.
func CreateEmptyIndMatrix(dim int) IndMatrix {
	values := make([]int, 0, dim)
	return IndMatrix{dim, values}
}

func (im *IndMatrix) AddRow(values []int) {
	if len(values) != im.Dim {
		log.Printf("[AddRow] Expected %d values, got %d \n.", im.Dim, len(values))
	}

	for _, e := range values[0:im.Dim] {
		im.Values = append(im.Values, e)
	}
}

// Returns i-th row of IndMatrix.
func (im *IndMatrix) GetRowValues(i int) []int {
	return make([]int, 0)
}

// Returns i-th col of IndMatrix.
func (im *IndMatrix) GetColValues(i int) []int {
	return make([]int, 0)
}
