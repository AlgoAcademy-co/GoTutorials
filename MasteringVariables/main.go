package main

import (
  "fmt"
)

//Constants are immuatable values that are known at compile time.
const pi = "3.14"

func main () {
  //Example of how variables can be declared 
  var name string

  //Strings are sequences of characters like the famous "Hello World", or "Algo Academy"
  var channelName string = "Algo Academy"
  
  //Integers are whole numbers, like 43, 99, or -10
  var age int
  age = 28

  //Booleans are either true or false 
  happy, success := true, false

  // Arrays have a fixed size and contain elements of the same type
  dogNames := [2]string {"Kylo", "Rey"}

  // Slices are dynamic arrays with a flexible size
  programmingLanguages := []string {"Go", "Python", "Java", "C"}

  // Maps are key-value pairs where  keys and values can be of any type
  scores := map[string] int {
    "Alice": 85,
    "Bob": 72,
    "Charlie": 98,
  }

  //Structs are custom types that group fields of different types
  type Person struct {
    Name string
    Age int
  }
  person1 := Person{Name: "Garrett", Age: 303}

  //Pointers store the memory address of a variable 
  var x int = -10
  ptr := &x

  //Type inference allows Go to determine the type of the variable based on its values
  number := 97
  message := "Hello World!"

  //Multiple assignments can be used to assigns values to multiple variables in one line
  var l, m, n = 8, 9, 10 

// Zero values are the default values assigned to variables when they are declared.
	var i int
	var f float64
	var b bool
	var s string


  //Printing the values of the variables
  fmt.Println("The name of my channel is:" + channelName)
	fmt.Printf("I am %v years old. \n", age)
	fmt.Println("Is there a reason to be happy today?", happy)
	fmt.Println("Will this channel succeed?", success)
	fmt.Printf("The name of my dogs are %v and %v", dogNames[0], dogNames[1])
	fmt.Println("Some popular programming languages are ", programmingLanguages[0], ",", programmingLanguages[1], ", and", programmingLanguages[2])
	fmt.Println("Alice's score is", scores["Alice"])
	fmt.Println("Name:", person1.Name, "Age:", person1.Age)
  fmt.Println("Value:", *ptr, "Address:", ptr)
	fmt.Println("Number:", number, "Message:", message)
	fmt.Println("Values:", l, m, n)
	fmt.Printf("Zero values: %v %v %v %q\n", i, f, b, s)
}
