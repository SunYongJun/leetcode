package transpose

import (
	"fmt"
	"testing"
)

func TestTranspose(t *testing.T) {
	fmt.Println(  Transpose([][]int{{1,2,3},{4,5,6},{7,8,9}}) )
}
