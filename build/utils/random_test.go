package utils

import (
	"fmt"
	"testing"
)

func TestRandomVideoNo(t *testing.T) {
	str := RandomVideoNo(6)
	fmt.Printf("rand str:%v\n", str)
}