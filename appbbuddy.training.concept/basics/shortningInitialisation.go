// when var are of same type they can be initialized together

package basics

import "fmt"

func add(x, y int) int {
	return x + y
}

func short() {
	fmt.Println(add(42, 13))
}
