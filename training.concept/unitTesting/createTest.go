package calculator

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// {"valid-data", 100.0, 10.0, 110.0, false},
// {"valid-data", 500.0, 10.0, 510.0, false},

func Genrate() {
	var aa float32
	var bb float32
	var cc float32
	var d []string

	for i := 0; i < 1000; i++ {
		aa = rand.Float32() * 100
		bb = rand.Float32() * 100
		cc = aa / bb
		output := fmt.Sprintf("{\"valid-data\", %f, %f, %f, false },", aa, bb, cc)
		fmt.Print(output)
		d = append(d, output)
		rand.NewSource(time.Now().UnixNano())
	}

	f, err := os.Create("lines")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
