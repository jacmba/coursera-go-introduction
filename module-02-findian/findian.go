/**
 * Module 02 - Exercise 2: Find 'IAN'
 *
 * Author: Jacinto Mba <jacmba(at)gmail(dot)com>
 *
 * https://github.com/jacmba/coursera-go-introduction
 *
 * Write a program which prompts the user to enter a string.
 * The program searches through the entered string for the characters
 * ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string
 * starts with the character ‘i’, ends with the character ‘n’, and contains
 * the character ‘a’. The program should print “Not Found!” otherwise.
 * The program should not be case-sensitive, so it does not matter if the
 * characters are upper-case or lower-case.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * find 'i-a-n' sequence inside given string
 */
func find(s string) bool {
	sl := strings.ToLower(s)
	return strings.HasPrefix(sl, "i") && strings.Contains(sl, "a") &&
		strings.HasSuffix(sl, "n")
}

func main() {
	var input string
	fmt.Println("Please enter a string:")

	// Using buffered IO to allow multiple words
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		panic("Input scanning failed!")
	}

	input = scanner.Text()

	found := find(input)
	if found {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
