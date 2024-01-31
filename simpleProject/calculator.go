package main

import (
	"os"
	"fmt"
)

// implicit type declaration by only declaring last parameter
func addValues(x , y int) int{
	return x + y
}


func multiplyValues(x , y int) int{
	return x * y
}

// implement function closer that calculates power in order to maintain "local variable" 
func pow() func(base, exponent int) int{
	output := 1
	return func (base, exponent int) int{
		for i := 1; i <= exponent; i++{
			output = output * base
		}
		return output

	}
}



func getInput() (int, int, string){
	fmt.Println("Welcome to a quick calculator demonstration!")
	
	fmt.Print("Please enter you're first integer A: ")
	var intA int
	_, err := fmt.Scanf("%d",&intA)
	fmt.Scanln()
	
	fmt.Print("Please enter you're first integer B: ")
	var intB int
	_, err = fmt.Scanf("%d",&intB)
	fmt.Scanln()
	
	// fmt.Printf("%d, %d \n", intA, intB)
	fmt.Print("Please enter your desired operation [add, multiply, pow]: ")
	var input string
	_, err = fmt.Scan(&input)
	fmt.Scanln()
	
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return intA, intB, input
}

func main(){
	intA, intB, operation := getInput()
	switch operation{
	case "add":
		result := addValues(intA, intB)
		fmt.Printf("%d + %d is %d", intA, intB, result)
	case "multiply":
		result := multiplyValues(intA, intB)
		fmt.Printf("%d * %d is %d", intA, intB, result)
	case "pow":
		exponential := pow()
		result := exponential(intA, intB)
		fmt.Printf("%d ^ %d is %d", intA, intB, result)
	}

}

