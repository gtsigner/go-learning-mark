package main;

import (
    "fmt"
)

/*Hello World*/
func main() {
	
	//new 数组 是得到一个指向数组的指针

	//arr
	var arr1 [2]int
	var arr2=[...]int{1,2,3,5,6,7,8}

	fmt.Println(arr1)

	//slice
	var s1 []int
	fmt.Println(s1)

	//5-所有的元素
	s2:=arr2[5:]

	//初始化一个容量5,申请一篇连续的内存地址空间。包含3个元素的数组
	s3:=make([]int,3,5)
	//s3[4]=10

	fmt.Println(s2)
	fmt.Println(s3)
	//获取容量
	fmt.Println(cap(s3),len(s3))


	
	

}
