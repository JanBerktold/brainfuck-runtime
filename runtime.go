package brain

import (
	"errors"
	"fmt"
)

// #include <stdio.h>
import "C"

type runtime struct {
	pointer int
	data    []int64
}

// doubles size of data array
func (r *runtime) growData() {
	newData := make([]int64, len(r.data)*2)
	copy(newData, r.data)
	r.data = newData
}

func new(size int64) *runtime {
	return &runtime{
		pointer: 0,
		data:    make([]int64, size),
	}
}

func getMatchingChar(code string, start, dir int) int {
	level := 0
	for i := start; i > 0 && i < len(code); i += dir {
		switch code[i] {
		case '[':
			level++
		case ']':
			level--
		}
		if level == 0 {
			return i
		}
	}
	return start
}

func Execute(code string) error {
	base := new(128)

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			base.pointer++
			if float32(base.pointer) > float32(len(base.data))*0.8 {
				base.growData()
			}
		case '<':
			base.pointer--
			if base.pointer < 0 {
				return errors.New("Pointer hit zero")
			}
		case '+':
			base.data[base.pointer]++
		case '-':
			base.data[base.pointer]--
		case '.':
			fmt.Print(string(base.data[base.pointer]))
		case '[':
			if base.data[base.pointer] == 0 {
				i = getMatchingChar(code, i, 1)
			}
		case ',':
			C.getchar()
		case ']':
			if base.data[base.pointer] > 0 {
				i = getMatchingChar(code, i, -1)
			}
		}
	}

	fmt.Println()
	return nil
}
