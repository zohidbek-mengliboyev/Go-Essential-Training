package main

import (
	"fmt"
)

func main() {
	// Same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Println("loons = %v (type %T)\n", loons, loons)

	// Length
	fmt.Println(len(loons)) // 3

	fmt.Println("------")
	// 0 indexing
	fmt.Println(loons[1]) // daffy

	fmt.Println("------")
	// slices
	fmt.Println(loons[1:]) // [daffy, taz]

	fmt.Println("------")
	// Single value range
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("------")
	// Double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}
	fmt.Println("------")
	// Double value range, ignore index by using
	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("------")
	// append
	loons = append(loons, "elmer")
	fmt.Println(loons) // [bugs, daffy, taz, elmer]

	fmt.Println("------")
	nums := []int{16, 8, 42, 4, 23, 15}
	max := nums[0] // Initialize max with first value
	// [1:] skips the first value
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)

}
