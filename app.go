package main

import "fmt"

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

}
