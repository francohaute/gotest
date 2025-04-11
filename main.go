package main

import "fmt"
import "time"

func main() {
	slice := make([]int, 100_000_000, 100_000_000)
	for i := range slice {
		slice[i] = i
	}
	start := time.Now()
	Filter(slice, func(i int) bool {
		return (i%2 == 0)
	})
	fmt.Println("Duration filter with: ", time.Since(start))
	start2 := time.Now()
	FilterWithoutAppend(slice, func(i int) bool {
		return (i%2 == 0)
	})
	fmt.Println("Duration filter without: ", time.Since(start2))
}
