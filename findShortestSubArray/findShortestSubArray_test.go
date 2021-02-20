package findShortestSubArray

import (
	"fmt"
	"testing"
)

func TestFindShortestSubArray(t *testing.T) {
	fmt.Println(FindShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(FindShortestSubArray([]int{1,2,2,3,1,4,2}))
}