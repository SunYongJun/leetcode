package isIsomorphic

import (
	"fmt"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	if IsIsomorphic("egg", "add") {
		t.Log("OK,通过")
	}

	if !IsIsomorphic("foo", "bar") {
		t.Log("OK")
	}

	if IsIsomorphic("paper", "title") {
		t.Log("OK")
	}

	fmt.Println(IsIsomorphic("badc","baba"))

}