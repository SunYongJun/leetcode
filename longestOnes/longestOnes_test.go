package longestOnes

import (
	"fmt"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	fmt.Println(LongestOnes([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3))
	fmt.Println(LongestOnes([]int{1,1,1,0,0,0,1,1,1,1,0}, 2))
	fmt.Println(LongestOnes([]int{1,0}, 2))
	fmt.Println(LongestOnes([]int{1}, 2))
	fmt.Println(LongestOnes([]int{0,0}, 0))
	fmt.Println(LongestOnes([]int{0,0,1,1,1,0,0}, 0))
}