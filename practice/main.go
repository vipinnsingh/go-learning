package main

import (
	"fmt"
	"strings"
)

func main() {

	// for i := 0; i < 5; i++ {
	// 	go func(i int) {
	// 		fmt.Println("i", i)
	// 	}(i)
	// }
	// time.Sleep(1 * time.Millisecond)
	input := "00010110"
	mask := "XX0XXX11"

	result := applyMask(input, mask)
	fmt.Printf("result: %v\n", result)
}

func applyMask(input string, m string) string {

	var result []string
	for i := 0; i < len(m); i++ {

		if string(m[i]) == "X" {
			result = append(result, string(input[i]))
		} else {
			result = append(result, string(m[i]))
		}
	}

	return strings.Join(result, "")

}
