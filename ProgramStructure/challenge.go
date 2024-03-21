package main

import (
  "fmt"
)

//Constants are immuatable values that are known at compile time.
const pi = "3.14"

func main () {

  //Strings - What's your newest favorite YouTube Channel that starts with an A and teaches you about Go programming?? Assign that value to the string below!
  var channelName string = ""
  
  //Define a variable named age, that is the integer data type

  //Booleans are either true or false 
  happy, success := true, false

  // Arrays have a fixed size and contain elements of the same type, change the array name and values it contains! Make sure to change the print statement below to accomodate
  arrayName := [2]string {"elementOne", "elementTwo"}

  // Slice - Populate this slice with 3 of your favorite programming languages
  // programmingLanguages := []string {}


  // Define a few variables while letting Go infer what their data type is. Ex. "number:= 97"

  //Multiple assignments can be used to assigns values to multiple variables in one line
  var l, m, n = 8, 9, 10 

// Zero values are the default values assigned to variables when they are declared.
	var i int
	var f float64
	var b bool
	var s string


  //Printing the values of the variables
  fmt.Println("The name of my channel is:" + channelName)
//	fmt.Printf("I am %v years old. \n", age)
	fmt.Println("Is there a reason to be happy today?", happy)
	fmt.Println("Will this channel succeed?", success)
	fmt.Printf("The items in my array are %v and %v", arrayName[0], arrayName[1])
	//fmt.Println("Some popular programming languages are ", programmingLanguages[0], ",", programmingLanguages[1], ", and", programmingLanguages[2])	
	fmt.Println("Values:", l, m, n)
	fmt.Printf("Zero values: %v %v %v %q\n", i, f, b, s)
}
