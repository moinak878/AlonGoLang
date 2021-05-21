// In a world full of variables , be my constant
package main

import "fmt"
// declaring variables and not using them is an error in go XD

func main(){

	// =============================String=================================

	var name string ="moinak"
	var name2 = "noob" // go automatically infers the type of variable 
	var name3 string // defualt value -> empty string 

	fmt.Println(name)
	fmt.Println(name2,name3)

	name="moinakv2.0"
	name3="piro"

	fmt.Println(name,name2,name3) // prints space separated values on screen

	
	name4 := "woah" // shorthand way of declaring variable
	fmt.Println(name4)

	// ====================================================================


	
	// =============================Integer=================================
	
	var age1 int = 20
	var age2 = 30
	age3 := 40

	fmt.Println(age1,age2,age3)
	// bits and memory
	// var age4 int8 = 215 //error : out of range
	// var age5 uint =123 //only positive numbers 

	// ====================================================================



	// =======================float========================================

	var scoreOne float32 = 15.7  // must specify whether 32 or 64 
	scoreTwo  := 13242341234134235.7 // float64 by default
	fmt.Println(scoreOne,scoreTwo)

}