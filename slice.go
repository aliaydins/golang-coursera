package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := make([]int, 0, 3)

	var scan string

	for {
		fmt.Print("Enter an integer :")
		fmt.Scan(&scan)

		exitCode := strings.ToLower(scan)

		if exitCode == "x" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		parseIntScanValue, _ := strconv.Atoi(scan)
		data = append(data, parseIntScanValue)
		sort.Ints(data)
		fmt.Println(data)
	}
}
