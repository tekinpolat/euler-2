package main

import (
	"fmt"
)

func main(){
	fmt.Println("Euler2 Starting...")
	total := 2

	fib_sequence := []int{1, 2}
	
	for index := 2; index <= 4000000; index++{
		value := fib_sequence[index - 1] + fib_sequence[index - 2]
		if value > 4000000{
			break
		}
		fib_sequence = append(fib_sequence, value)

		if( value % 2 == 0){
			total += value
		}
	} 
	fmt.Println("Total:", total)
}