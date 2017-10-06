package main

import (
	"fmt"
	"./lib"
	"math"
)

type(
	字符串 string
)
func main(){
	lib.Demo();
	fmt.Println("Hello\n")

	//默认值是0
	var a int
	var b int32
	var c int64
	var d bool
	var e float32
	//切片
	var f []int32
	var g string
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var str 字符串
	str="Hello 中文"
	fmt.Println(str)
	fmt.Println(math.MaxInt32)

	/*类型转换*/


	var myA int= 1
	var myB float32= float32(myA)
	myA=int(myB)

	print(myA)
	print(myB)

}
