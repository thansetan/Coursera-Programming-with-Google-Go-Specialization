package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either
create a new animal of one of the three types, or the user can request information about an animal
that he/she has already created. Each animal has a unique name, defined by the user.
Note that the user can define animals of a chosen type, but the types of animals are restricted
to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

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
Your program should accept one command at a time from the user, print out a response, and print out
a new prompt on a new line. Your program should continue in this loop forever. Every command from the user
must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal. The third string is
the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should process each
newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”.
The second string is the name of the animal. The third string is the name of the information requested about the animal,
either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically,
the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.

1. The Eat() method should print the animal’s food,
2. the Move() method should print the animal’s locomotion, and
3. the Speak() method should print the animal’s spoken sound.

Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak()
so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal,
create an object of the appropriate type. Your program should call the appropriate method
when the user issues a query command.
*/

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, locomotion, noise string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	food, locomotion, noise string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	food, locomotion, noise string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

var animals = make(map[string]Animal)

func main() {
	reader := bufio.NewReader(os.Stdin)
	validCommand := []string{"newanimal", "query"}
	validAnimal := []string{"cow", "bird", "snake"}
	validRequest := []string{"eat", "speak", "move"}
	for {
		var commandSlice []string
		fmt.Printf(">")
		commands, _ := reader.ReadString('\n')

		commandSlice = strings.Fields(strings.ToLower(commands))

		if len(commandSlice) != 3 {
			fmt.Println("invalid command, command must be 3 strings separated by a space: \"string1 string2 string3\"")
			continue
		}

		if !isIn(commandSlice[0], validCommand) {
			fmt.Printf("%s is not a valid command, valid commands are: %v\n", commandSlice[0], validCommand)
			continue
		}

		if commandSlice[0] == "newanimal" {
			if !isIn(commandSlice[2], validAnimal) {
				fmt.Printf("%s is not a valid animal type, valid animal types are: %v\n", commandSlice[2], validAnimal)
				continue
			}
		} else {
			if !isIn(commandSlice[2], validRequest) {
				fmt.Printf("%s is not a valid request, valid requests are: %v\n", commandSlice[2], validRequest)
				continue
			}
		}

		switch commandSlice[0] {
		case "newanimal":
			_, exist := findAnimalByName(commandSlice[1])
			if exist {
				fmt.Printf("animal named %s already exist, please choose another name\n", commandSlice[1])
				continue
			}
			createNewAnimal(commandSlice[1], commandSlice[2])
			fmt.Println("Created it!")
		case "query":
			animal, exist := findAnimalByName(commandSlice[1])
			if !exist {
				fmt.Printf("animal named %s does not exist!\n", commandSlice[1])
				continue
			}
			queryAnimal(animal, commandSlice[2])
		}
	}
}

func isIn(val string, values []string) bool {
	for _, value := range values {
		if val == value {
			return true
		}
	}
	return false
}

func createNewAnimal(name, animalType string) {
	switch animalType {
	case "cow":
		animals[name] = Animal(newCow())
	case "bird":
		animals[name] = Animal(newBird())
	case "snake":
		animals[name] = Animal(newSnake())
	}
}

func queryAnimal(animal Animal, request string) {
	switch request {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}

}

func findAnimalByName(name string) (Animal, bool) {
	animal, exist := animals[name]
	return animal, exist
}

func newCow() Cow {
	cow := Cow{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	return cow
}

func newBird() Bird {
	bird := Bird{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	return bird
}

func newSnake() Snake {
	snake := Snake{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
	return snake
}
