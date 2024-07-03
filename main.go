package main

import "fmt"

// import "fmt"


func main() {


	var a int = 12+12+1212143564754321456
    fmt.Println(a)
	// var (
	// 	a = 7.0
	// 	b = 3.0
	// )

	// fmt.Println(a/b)


// var c complex64 = complex(1,2) // 1+2i
// var d complex128 = complex(3,4) // 3+4i

// fmt.Println(c)
// fmt.Println(d)




// var x int = 5
// fmt.Println(x)

// x := 5
// fmt.Println(x)

// var (
// 	x = 5
// )
// fmt.Println(x)



// var x byte = 'B'
// var q byte = 'L'
// var z byte = 'U'
// var n byte = 'E'
// fmt.Println(x,q,z,n)

// var x rune = '9'
// var q rune = 'b'
// var z rune = 'A'
// fmt.Println(x,q,z)


// var e byte = 'a'
// var r rune = '♥'
// fmt.Println(e,r)





// var complexNumber1, complexNumber2, complexNumber3 complex64

// // initializing the variable using complex number init syntax
// complexNumber1 = 3 + 3i
// complexNumber2 = 2 + 5i

// fmt.Println("The First Complex Number is", complexNumber1)
// fmt.Println("The Second Complex Number is", complexNumber2)

// // adding the complex numbers using + operator
// complexNumber3 = complexNumber1 + complexNumber2

// fmt.Println("Printing the addition of two complex numbers by initializing the variable using complex number init syntax.")

// // printing the complex number after addition
// fmt.Println(complexNumber1, "+", complexNumber2, "=", complexNumber3)




// var complexNumber1, complexNumber2, complexNumber3 complex64
   
//    // initializing the variable using the constructor
//    complexNumber1 = complex(5, 4)
//    complexNumber2 = complex(6, 3)
   
//    fmt.Println("The First Complex Number is", complexNumber1)
//    fmt.Println("The Second Complex Number is", complexNumber2)
   
//    // adding the complex numbers using + operator
//    complexNumber3 = complexNumber1 + complexNumber2
   
//    fmt.Println("Printing the addition of two complex numbers by initializing the variable using the constructor.")

//    // printing the complex number after addition
//    fmt.Println(complexNumber1, "+", complexNumber2, "=", complexNumber3)

// str1 := "Hello "
// str2 := "World!"
// result := str1 + str2
// fmt.Println(result) // Output: Hello, World!



// str1 := "Hello"
// str2 := "World"
// fmt.Println(str1 == str2) // Output: false
// fmt.Println(str1 < str2)  // Output: true

// var b byte = 'a' // ASCII value of 'a' is 97
// fmt.Println(b) // Output: 97


// var r rune = '世' // Unicode value for '世' is 19990
// fmt.Println(r) // Output: 19990
// fmt.Printf("%c\n", r) // Output: 世


// arr := [5]int{1, 2, 3, 4, 5}

// for i := 0; i < len(arr); i++ {
//     fmt.Println(arr[i])
// }


// slice := []int{1, 2, 3}

// slice = append(slice, 4, 5)

// fmt.Println(slice)


// slice := []int{1, 2, 3, 4, 5}
// slice[2] = 10
// fmt.Println(slice)


// slice := []int{1, 2, 3, 4, 5}

// indexToDelete := 1

// slice = append(slice[:indexToDelete], slice[indexToDelete+1:]...)

// fmt.Println(slice)


// originalSlice := []int{1, 2, 3, 4, 5, 6, 7}
// newSlice := originalSlice[2:4] // [2, 3, 4]

// // newSlice = append(newSlice, 1,2)

// fmt.Println(newSlice)



// sourceSlice := []int{1, 2, 3}
// targetSlice := make([]int, 10)
// copy(targetSlice, sourceSlice)


// fmt.Println(targetSlice)
// fmt.Println(sourceSlice)



// var m1 map[string]int
//     // Initialize the map
//     m1 = make(map[string]int)

//     // Adding key-value pairs
//     m1["dog"] = 4
//     m1["cat"] = 9
//     fmt.Println("Initial map:", m1) // Output: Initial map: map[dog:4 cat:9]

//     // Checking if a key exists
//     value, ok := m1["dog"]
//     if ok {
//         fmt.Println("Key 'dog' exists with value:", value) // Output: Key 'dog' exists with value: 4
//     } else {
//         fmt.Println("Key 'dog' does not exist")
//     }



// var m1 map[string]int
// // Initialize the map
// m1 = make(map[string]int)

// // Sample data: list of strings
// data := []string{"apple", "banana", "apple", "orange", "banana", "apple"}

// // Counting occurrences of each string
// for _, item := range data {
// 	m1[item]++
// }

// // Printing the map with counts
// fmt.Println("String counts:", m1) // Output: String counts: map[apple:3 banana:2 orange:1]


// fmt.Println("Count for 'apple':", m1["apple"]) // Output: Count for 'apple': 3
// fmt.Println("Count for 'orange':", m1["orange"]) // Output: Count for 'orange': 1
// fmt.Println("Count for 'grape':", m1["grape"]) // Output: Count for 'grape': 0


// type Person struct {
//     Name string
//     Age  int
// }

// p1 := Person{"Alice",30}
// p2 := Person{Name:"Alice",Age: 30}


// fmt.Println(p1)
// fmt.Println(p2)
}


