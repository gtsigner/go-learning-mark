package main

import (
	"fmt"
)

func main() {

	/*Containue循环*/
HELLO:
	for i := 1; i < 10000; i++ {
		for l := 10; l < 100; l++ {
			fmt.Println("%d---%d",i,l)
			//GOTO
			continue HELLO
		}
	}

}
