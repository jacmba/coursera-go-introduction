/* Module 04- Exercise 1: Make JSOn
 *
 * Author: Jacinto Mba <jacmba(at)gmail(dot)com>
 *
 * https://github.com/jacmba/coursera-go-introduction
 *
 * Write a program which prompts the user to first enter a name, and then enter an address.
 * Your program should create a map and add the name and address to the map using the keys
 * “name” and “address”, respectively. Your program should use Marshal() to create a JSON
 * object from the map, and then your program should print the JSON object.
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/**
 * makejson - Get a key-value map and return corresponding JSON string reresentation
 */
func makejson(input map[string]string) string {
	j, _ := json.Marshal(input)
	return string(j)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make(map[string]string)

	fmt.Print("Enter name: ")
	if !scanner.Scan() {
		panic("Big trouble!")
	}
	data["name"] = scanner.Text()

	fmt.Print("Enter address: ")
	if !scanner.Scan() {
		panic("Big trouble!")
	}
	data["address"] = scanner.Text()

	j := makejson(data)
	fmt.Println(j)
}
