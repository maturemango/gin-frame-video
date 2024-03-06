package utils

import (
	"fmt"
	"log"
)

type Action func() error

type Actions struct {
	Name string
	Func Action
}

type FuncAction []Actions

func (f FuncAction) Do() {
	for _, a := range f {
		if err := a.Func(); err != nil {
			panic(fmt.Errorf("%s failed, error %s", a.Name, err))
		}
		log.Printf("%s success\n", a.Name)
	}
}