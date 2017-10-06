package main

import (
	"fmt"
)

const a int = 1
const b = 100
const c = 'a'
const d = "A"


const (
	e=1
	f=2
)

/*Error const在编译的时候进行，var运行时*/
// var str="hello"
// const (
// 	someLen=len(str)
// )


//是常量组中的记数器iota
const (
	one = iota
	two
	three
)

func main(){


	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)

	fmt.Println(e)
	fmt.Println(f)

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	
}