// similar to java for loop but no ()

package flow

import "fmt"

func forloop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}
