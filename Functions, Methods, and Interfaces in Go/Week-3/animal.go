package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
1) the food that it eats,
2) its method of locomotion, and
3) the sound it makes when it speaks.

The following table contains the three animals and their associated data which should be hard-coded into your program.

+--------+------------+-------------------+--------------+
| Animal | Food eaten | Locomotion method | Spoken sound |
+--------+------------+-------------------+--------------+
| Cow    | Grass      | Walk              | Moo          |
+--------+------------+-------------------+--------------+
| Bird   | Worms      | Fly               | Peep         |
+--------+------------+-------------------+--------------+
| Snak   | Mice       | Slither           | Hsss         |
+--------+------------+-------------------+--------------+

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request,
and prints out a new prompt. Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings. The first string is the name of an animal,
either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal,
either “eat”, “move”, or “speak”. Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which
is a struct containing three fields:food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method
when the user makes a request.
*/

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	animalName := []string{"cow", "bird", "snake"}
	allowedReq := []string{"eat", "move", "speak"}
	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	for {
		var (
			userRequestStr string
			userRequest    []string
		)

		fmt.Printf(">")
		userRequestStr, _ = reader.ReadString('\n')

		userRequest = strings.Fields(strings.ToLower(userRequestStr))

		if len(userRequest) != 2 {
			fmt.Println("invalid request, request must be in form: [animal_name <space> something_you_want_to_get]")
			continue
		}

		if !isIn(userRequest[0], animalName) {
			fmt.Printf("invalid animal name, available name: %v\n", animalName)
			continue
		}

		if !isIn(userRequest[1], allowedReq) {
			fmt.Printf("invalid request name, available request: %v\n", allowedReq)
			continue
		}

		switch userRequest[0] {
		case "cow":
			getRequestedMethod(cow, userRequest[1])
		case "bird":
			getRequestedMethod(bird, userRequest[1])
		case "snake":
			getRequestedMethod(snake, userRequest[1])
		}
	}
}

func isIn(val string, values []string) bool {
	for _, value := range values {
		if value == val {
			return true
		}
	}
	return false
}

func getRequestedMethod(a Animal, req string) {
	switch req {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}
