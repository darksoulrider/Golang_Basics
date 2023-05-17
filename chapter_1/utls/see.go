package utls

import (
	// "bytes"
	"fmt"
)

func See() {
	data := "I am good"
	strBytes := []byte(data)

	fmt.Println(strBytes)
	fmt.Println(string(strBytes[0]))
}
