// * for value & for pointer to adress
// Unlike C, Go has no pointer arithmetic.

package main

import "fmt"

func main() {
	PointerBasic()
	PointExample()

	a := 4
	b:= 5
	PassValueAndPointer(a,&b)
	fmt.Printf("a = {a}, b={b}")
}

func PointerBasic(){
	var myNum = 3
	var ptr *int //initialize pointer
	ptr = &myNum //point with pointer

	fmt.Println("value of ptr -",ptr )
	fmt.Println("value ptr is pointing towards -",*ptr)
}

func PointExample() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func PassValueAndPointer(value int,ptr *int){
	value = value + 1
	*ptr = *ptr +1
}