/**
 * Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
 * The program should be written as a loop.
 * Before entering the loop, the program should create an empty integer slice of size (length) 3.
 * During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
 * The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
 * The slice must grow in size to accommodate any number of integers which the user decides to enter.
 * The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
 * 30 20 15 25 5 10
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Create a new slice with a len/cap 3
	var intArray = make([]int64, 0, 3)

	scanner := bufio.NewScanner(os.Stdin)

	exiting := false
	for {
		fmt.Print("Please enter an integer: ")
		scanner.Scan()

		if scanner.Err() != nil {
			continue
		}

		line := scanner.Text()
		tokens := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })
		for _, token := range tokens {
			if strings.ToLower(token) == "x" {
				exiting = true
				break
			}
			intValue, err := strconv.ParseInt(token, 10, 64)
			if err != nil {
				continue
			}
			intArray = sortedAdd(intArray, intValue)
		}

		// Display contents
		display(intArray)

		// Break out of infinite loop?
		if exiting {
			break
		}
	}
}

func display(list []int64) {
	fmt.Print("[")
	listSize := len(list) - 1
	for i, v := range list {
		fmt.Printf("%d", v)
		if i < listSize {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func sortedAdd(list []int64, value int64) []int64 {
	if len(list) == 0 {
		return append(list, value)
	}

	if len(list) == 1 {
		if list[0] < value {
			return append(list, value)
		} else {
			return append([]int64{value}, list...)
		}
	}
	index, _ := slices.BinarySearch(list, value)
	list = slices.Grow(list, 1)
	return slices.Insert(list, index, value)
}
