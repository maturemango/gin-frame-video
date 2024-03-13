package utils

import (
	"fmt"
	"testing"
)

func TestRandomVideoNo(t *testing.T) {
	str := RandomNo(6)
	fmt.Printf("rand str:%v\n", str)
}