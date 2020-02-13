/**
 * Module 03 - Exercise 1: Slice
 *
 * Author: Jacinto Mba <jacmba(at)gmail(dot)com>
 *
 * https://github.com/jacmba/coursera-go-introduction
 *
 * Write a program which prompts the user to enter integers and stores the
 * integers in a sorted slice. The program should be written as a loop.
 * Before entering the loop, the program should create an empty integer slice
 * of size (length) 3. During each pass through the loop, the program prompts
 * the user to enter an integer to be added to the slice. The program adds the
 * integer to the slice, sorts the slice, and prints the contents of the slice
 * in sorted order. The slice must grow in size to accommodate any number of
 * integers which the user decides to enter. The program should only quit
 * (exiting the loop) when the user enters the character ‘X’ instead of an integer.
 */

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)
	n := 0
	quit := false

	for !quit {
		var input string
		fmt.Println("Please enter an integer (X for exit):")
		_, err := fmt.Scan(&input)
		if err != nil {
			panic(err)
		}

		if input == "x" || input == "X" {
			quit = true
		} else {
			num, err := strconv.Atoi(input)

			if err != nil {
				panic(err)
			}

			if n < 3 {
				sli[n] = num
			} else {
				sli = append(sli, num)
			}

			sorted := make([]int, len(sli))
			copy(sorted, sli)
			sort.Ints(sorted)

			fmt.Println(sorted)
			n++
		}
	}

	fmt.Println("Bye!")
}
