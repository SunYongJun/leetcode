package fairCandySwap

import (
	"fmt"
	"testing"
)

func TestFairCandySwap(t *testing.T) {
	fmt.Println(FairCandySwap([]int{1,1}, []int{2,2}))

	fmt.Println(FairCandySwap([]int{1,2}, []int{2,3}))

	fmt.Println(FairCandySwap([]int{2}, []int{1,3}))

	fmt.Println(FairCandySwap([]int{1,2,5}, []int{2,4}))

}