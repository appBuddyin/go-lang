package calculator

import (
	"log"
	"strings"
)

func Cal() {
	a := strings.Split("/1/*/3", "/")
	log.Println(a[:], len(a))
	// router := gin.Default()
	// router.GET("/*path", func(c *gin.Context) {
	// 	path := c.Param("path")
	// 	// a := c.Param("a")
	// 	// o := c.Param("o")
	// 	// b := c.Param("b")
	// 	// d := reflect.ValueOf(b).Kind()
	// 	fmt.Printf("type = %T", path)

	// 	message := ":" + path + ":"

	// 	log.Printf("path type = %T  ", path)
	// 	// message := d
	// 	c.String(http.StatusOK, message)
	// })

	// router.Run(":8080")
}
