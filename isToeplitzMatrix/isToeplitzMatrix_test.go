package isToeplitzMatrix

import (
	"fmt"
	"testing"
)

func TestIsToeplitzMatrix(t *testing.T) {
	fmt.Println(IsToeplitzMatrix([][]int{{1,2,3,4},{5,1,2,3},{9,5,1,2}}))
	fmt.Println(IsToeplitzMatrix([][]int{{1,2},{2,2}}))
}
