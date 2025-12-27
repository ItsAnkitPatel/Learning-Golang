package main

import "fmt"

// golang only have single loop, for
func main(){
	
	// while loop
	i := 1
	for i <=3 {
		fmt.Println(i)
		i += 1
	}
	// infinite loop
	// for {
	// 	println("1")
	// }

	// classic for loop
	// for i:=0; i< 3;i++{
	// 	fmt.Println(i)
	// }

	//in go 1.22, 'range' another way using of using for loop came
	for i := range 3 {
		// 3 will be excluded
		fmt.Println(i)
	}

}
