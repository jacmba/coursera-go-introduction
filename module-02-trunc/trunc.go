/**
 * Module 02 - Exercise 1: Truncate input number
 *
 * Author: Jacinto Mba <jacmba(at)gmail(dot)com>
 *
 * https://github.com/jacmba/coursera-go-introduction
 *
 * Write a program which prompts the user to enter a floating point number and
 * prints the integer which is a truncated version of the floating point number
 * that was entered. Truncation is the process of removing the digits to the
 * right of the decimal place.
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Function to truncate decimal positions in a nueric string and return
 * the converted integer number
 */
func truncate(s string) int {
	sTrunc := strings.Split(s, ".")[0]
	res, _ := strconv.Atoi(sTrunc) // Let's ignore the error and just return 0 value
	return res
}

func main() {
	fmt.Println("Please enter a number:")
	var input string
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("There was an error with your entered number:", err)
		return
	}

	result := truncate(input)
	fmt.Println("Your truncated integer number is:", result)
}
