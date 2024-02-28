# Mastering Variables in Go

This repository contains a Go program that showcases different types of variables and how they are used in Go programming. Let's explore the various types of variables used in the program:

1. **Constants (`const`):** Constants are immutable values that are known at compile time. In this program, we have a constant `pi` with the value "3.14".

2. **Strings (`string`):** Strings are sequences of characters. The `channelName` variable is a string that stores the name of a channel, "Algo Academy".

3. **Integers (`int`):** Integers are whole numbers. The `age` variable is an integer that stores a person's age, which is initialized to 28.

4. **Booleans (`bool`):** Booleans are either true or false. The `happy` and `success` variables are boolean variables initialized to true and false, respectively.

5. **Arrays (`[]string`):** Arrays have a fixed size and contain elements of the same type. The `fruits` array contains three strings: "apple", "banana", and "cherry".

6. **Slices (`[]string`):** Slices are dynamic arrays with a flexible size. The `colors` slice contains three strings: "red", "green", and "blue".

7. **Maps (`map[string]int`):** Maps are key-value pairs where keys and values can be of any type. The `scores` map stores the scores of three people: "Alice", "Bob", and "Charlie".

8. **Structs (`struct`):** Structs are custom types that group fields of different types. The `Person` struct has two fields: `Name` (string) and `Age` (int), used to store a person's name and age.

9. **Pointers (`*int`):** Pointers store the memory address of a variable. The `ptr` pointer points to the memory address of the `x` variable.

10. **Type Inference:** Go can determine the type of a variable based on its value. For example, the `number` variable is inferred as an integer with the value 42.

11. **Multiple Assignment:** Go allows assigning values to multiple variables in a single line. For example, `var a, b, c = 5, 7, 9`.

12. **Zero Values:** Variables declared without an explicit initial value are set to their zero value. For example, `var i int` is initialized to 0.

Feel free to explore the program to see these variables in action and how they are used in Go programming!

