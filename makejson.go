package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var m map[string]string

func main() {
	m = make(map[string]string)

	in := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter name : ")
	in.Scan()
	m["name"] = in.Text()
	fmt.Printf("\nEnter address : ")
	in.Scan()
	m["address"] = in.Text()

	jsonData, err := json.Marshal(m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonData))

}
