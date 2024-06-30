package main

import "fmt"


func main() {

	var (
		a = 7.0
		b = 3.0
	)

	fmt.Println(a/b)


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

var x rune = '9'
var q rune = 'b'
var z rune = 'A'
fmt.Println(x,q,z)


var e byte = 'a'
var r rune = '♥'
fmt.Println(e,r)



}