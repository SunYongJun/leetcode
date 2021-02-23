package maxSatisfied

import (
	"fmt"
	"testing"
)

func TestMaxSatisfied(t *testing.T) {
	fmt.Println(MaxSatisfied([]int{1,0,1,2,1,1,7,5}, []int{0,1,0,1,0,1,0,1}, 3))
	fmt.Println(MaxSatisfied([]int{4,10,10}, []int{1,1,0}, 2))
}
