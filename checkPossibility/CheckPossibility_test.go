package checkPossibility

import (
	"fmt"
	"testing"
)

func TestCheckPossibility(t *testing.T) {
	fmt.Println(CheckPossibility([]int{4,2,3}))
	fmt.Println(CheckPossibility([]int{4,2,1}))
	fmt.Println(CheckPossibility([]int{5,7,1,8}))
	fmt.Println(CheckPossibility([]int{5}))
	fmt.Println(CheckPossibility([]int{1,2,3,5,4}))
}
