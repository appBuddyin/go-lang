// can return more then a result like python
package basics

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func multiple() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
