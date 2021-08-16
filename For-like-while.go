// The init and post statements are optional.
// work like while
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite for
	// for {
	// }
		
}
