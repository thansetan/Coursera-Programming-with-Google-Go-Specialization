package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”,
respectively. Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
*/

func main() {
	personMap := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a name: ")
	name, _ := reader.ReadString('\n')
	personMap["name"] = strings.TrimSpace(name)

	fmt.Printf("Enter an address: ")
	address, _ := reader.ReadString('\n')
	personMap["address"] = strings.TrimSpace(address)

	b, _ := json.Marshal(personMap)
	fmt.Println(string(b))
}
