package main

import (
	"fmt"
)


/*!!! Go中是将数组直接进行Copy的*/
func writeA(a [6]int) {
	a[2]=100
}

/*Hello World*/
func main() {
	/*...直接进行赋值*/
	a := [...]int{1, 2, 2, 3, 4, 5}
	c := [...]int{19: 100}
	b := [10]int{9: 100}


	writeA(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	println(c[1])
	println(b[1])

	/*指向数组的指针*/

	var p *[20]int = &c
	println(p)

	/*数组的指针，和指向数组的指针*/

	dd:= [2][3] int{
		{1,2,3},
		{3,4,5}	}

	fmt.Println(dd)

}

