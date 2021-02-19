package main

import (
	"fmt"
)

func main (){

	var operation string
	var number1,number2 int
	fmt.Print("Enter the First number: ")
	fmt.Scanln(&number1)
	fmt.Print("Enter the Second Number: ")
	fmt.Scanln(&number2)
	fmt.Print("Enter the Operator( *,/,+,-): ")
	fmt.Scanln(&operation)

	calculate(number1,number2,operation)


}

func calculate(number1,number2 int,operation string){

	var output int

	if(operation == "+"){
		output = number1 + number2
		fmt.Println(number1,operation,number2 ,"=", output)
	} else if( operation == "-"){
		output = number1 - number2
		fmt.Println(number1,operation,number2 ,"=", output)
	}else if( operation == "/"){
		output = number1 / number2
		fmt.Println(number1,operation,number2 ,"=", output)
	}else if( operation == "*"){
		output = number1 * number2
	  fmt.Println(number1,operation,number2 ,"=", output)
	} else{
		validate(number1,number2,operation)
	}
}

func validate(num1, num2 int, operation string){
	if(num1 == 0 || num2==0 || operation == ""){
		fmt.Println("Cannot not calculate Empty Input")
	}else{
		calculate(num1,num2,operation)
	}
}