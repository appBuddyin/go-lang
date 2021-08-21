// The init and post statements are optional.
// work like while

package flow

import (
	"fmt"
)
func forLike() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite for
	// for {
	// }
}
