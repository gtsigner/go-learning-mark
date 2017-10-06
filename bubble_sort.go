package main

import (
	"fmt"
)

func main() {

	a := [...]int{1, 123, 123, 12, 31, 3, 123, 2, 4, 4, 45, 5, 6, 7, 7, 78, 2, 56, 6}

	count := len(a)
	println(count)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if a[i]-a[j] > 0 {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}
	}

	fmt.Println(a)

}
