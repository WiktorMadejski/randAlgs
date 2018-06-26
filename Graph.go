package randAlgs

// IndMatrix represents a graph as incidence matrix.
type IndMatrix []int

// Returns i-th row of IndMatrix.
func (im *IndMatrix) GetRow(i int) []int {
	return make([]int, 0)
}

// Returns i-th col of IndMatrix.
func (im *IndMatrix) GetCol(i int) []int {
	return make([]int, 0)
}

func CreateEmptyIndMatrix() IndMatrix {

}
