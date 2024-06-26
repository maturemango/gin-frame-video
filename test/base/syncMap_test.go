package base

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Map可提供一个安全并发的映射，特别是数据的读删操作性能极佳

var s = &sync.Map{}

func Test_CompareAndSwap(t *testing.T) {
	s.Store("123", "123")
	s.CompareAndSwap("123", "123", "1")
	value, _ := s.Load("123")
	fmt.Printf("%s\n", value)
}

func Test_CompareAndDelete(t *testing.T) {
	deleted := s.CompareAndDelete("123", "1")
	if deleted {
		fmt.Printf("one delete success")
	} else {
		s.Store("123", "12")
		deleted := s.CompareAndDelete("123", "1")
		if deleted {
			fmt.Printf("two delete success")
		} else {
			s.Store("123", "1")
			deleted := s.CompareAndDelete("123", "1")
			if deleted {
				fmt.Printf("three delete success")
			}
		}
	}
}

func Test_Range(t *testing.T) {
	for i := 0; i < 5; i++ {
		s.Store(i, i)
	}
	s.Range(func(key any, value any) bool {
		keyType := key.(int)
		if keyType > 5 {
			return false
		}
		v, _ := s.Load(key)
		fmt.Printf("value is %v\n", v)
		return true
	})
}