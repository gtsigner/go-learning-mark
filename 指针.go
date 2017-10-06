package main


import (
	"fmt"
)


func main(){

	a:=1


	var p *int= &a

	/*指针不能直接进行指针预算，可以变相运算*/

	//只能是单独的语句，只能放置在右边
	a++

	fmt.Println(p)
	fmt.Println(*p)
}