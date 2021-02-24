package flipAndInvertImage

import (
	"fmt"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	fmt.Println(FlipAndInvertImage([][]int{{1,1,0},{1,0,1},{0,0,0}}))
	fmt.Println( FlipAndInvertImage([][]int{{1,1,0,0},{1,0,0,1},{0,1,1,1},{1,0,1,0}}) )
}