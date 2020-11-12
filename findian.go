package main

import (
	"fmt"
	"strings"
)

func main() {
	var data string

	fmt.Printf("Enter a string : ")
	fmt.Scan(&data)

	lowerData := strings.ToLower(data)

	if strings.HasPrefix(lowerData, "i") && strings.Contains(lowerData, "a") && strings.HasSuffix(lowerData, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}

}
