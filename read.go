package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func charSize(data string) string {
	dataSize := []rune(data)
	return string(dataSize[0:20])
}
func main() {

	var person Person
	var fileName string
	newSlice := make([]Person, 0)

	fmt.Printf("Enter file name : ")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	in := bufio.NewScanner(file)
	for in.Scan() {
		s := strings.Split(in.Text(), " ")
		if len(s[0]) > 20 {
			s[0] = charSize(s[0])
		}
		if len(s[1]) > 20 {
			s[1] = charSize(s[1])
		}
		person.fname, person.lname = s[0], s[1]
		newSlice = append(newSlice, person)

	}
	for _, item := range newSlice {
		fmt.Println(item.fname, item.lname)
	}

}
