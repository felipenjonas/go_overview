package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter numbers separated by spaces: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fields := strings.Fields(input)
	sum := 0

	for _, field := range fields {
		n, err := strconv.Atoi(field)
		if err != nil {
			fmt.Printf("Skipping invalid number: %s\n", field)
			continue
		}
		sum += n
	}

	fmt.Printf("Sum: %d\n", sum)
}
