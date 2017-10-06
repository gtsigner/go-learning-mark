package main


import (
	"fmt"
)

const (
	B float64 =1<<(iota*10)
	KB
	MB
	GB
	TB
	EB
	ZB
)


func main(){


	
	//1
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(ZB)



}