package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file
has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name,
and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively
read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs
and print the first and last names found in each struct.
*/

type Name struct {
	fName, lName string
}

func main() {
	var nameSlice []Name
	var filename string

	fmt.Printf("Enter your filename: ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error reading file: %s\n", err.Error())
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		names := strings.Split(line, " ")
		if len(names) != 2 {
			fmt.Println("each line should have first name and last name separated by space (so there are 2 words separated by space)")
			os.Exit(1)
		}
		nameSlice = append(nameSlice, Name{
			fName: names[0],
			lName: names[1],
		})
	}

	if len(nameSlice) == 0 {
		fmt.Println("empty name slice")
	}

	for i, name := range nameSlice {
		fmt.Printf("%d. First name  : %s\n   Last name   : %s\n", i+1, name.fName, name.lName)
		fmt.Println("===================================")
	}
}
