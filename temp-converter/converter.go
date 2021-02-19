package main

import (
	"fmt"
)


func main (){
	var temp float64
	var measurement string
	

	fmt.Print("What are you converting to( C = Celsius , F = Franhiet, K = Kelvin): ")
	fmt.Scanln(&measurement)
	fmt.Print("What is the Temperature Degree: ")
	fmt.Scanln(&temp)

	convert(temp, measurement)
}

func convert (temp float64, measurement string){

	 var answer float64
	if(measurement == "C"){
		answer = temp * 9/5 + 32
		fmt.Println(answer)
	} else if(measurement == "F"){
		answer = (temp - 32) * 5/9
		fmt.Println(answer)
	} else if(measurement == "K"){
		answer = temp + 273.15 
		fmt.Println(answer)
	} else{
		validate(temp, measurement)
	}
}

func validate (temp float64, measurement string){
	if(temp == 0 || measurement == ""){
		fmt.Println("Numbers and measurement is needed")
	}
}

