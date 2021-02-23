package matrixReshape

import (
	"fmt"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	fmt.Println(MatrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(MatrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
	fmt.Println(MatrixReshape([][]int{{1, 2}, {3, 4},{5, 6}}, 2, 3))
}