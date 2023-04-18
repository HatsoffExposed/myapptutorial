package main

import (
	"errors"
	"fmt"
	geometry "forgames/dir2"
)

func sum(nums ...int) {
	fmt.Print(nums, "")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func zeroval(ival int) {
	ival = 0
	// fmt.Println("ival:", ival)
}

func zerpptr(iptr *int) {
	*iptr = 0
}

// Using built-in interface of type error
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// creating custom types as errors
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	i := 1
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zerpptr(&i)
	fmt.Println("zerpptr:", i)

	fmt.Println("pointer:", &i)

	r := geometry.Rect{Width: 3, Height: 4}
	c := geometry.Circle{Radius: 5}

	geometry.Measure(r)
	geometry.Measure(c)
	// arr := []int{7, 42}
	//Testing out the error returning functions
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	// fmt.Println(arr)
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
