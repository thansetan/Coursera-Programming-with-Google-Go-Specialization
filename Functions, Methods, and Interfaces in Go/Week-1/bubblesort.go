package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers
as an argument and returns nothing. The BubbleSort() function should modify the slice so that
the elements are in sorted order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

func main() {
	var (
		intInput   string
		sliceOfInt []int
	)
	fmt.Printf("please enter some space separated integers (max 10 integers, example: \"1 -1 3 99 1 5\"): ")
	reader := bufio.NewReader(os.Stdin)
	intInput, _ = reader.ReadString('\n')
	for _, numStr := range strings.Fields(intInput) {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			sliceOfInt = append(sliceOfInt, num)
		}
	}
	fmt.Printf("your input: %v\n", sliceOfInt)
	if len(sliceOfInt) > 10 {
		fmt.Printf("you are only allowed to enter max 10 integers, but you entered %d integers.\n", len(sliceOfInt))
		os.Exit(1)
	}
	BubbleSort(sliceOfInt)
	fmt.Printf("sorted input: %v\n", sliceOfInt)
}

func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

func Swap(nums []int, idx int) {
	nums[idx], nums[idx+1] = nums[idx+1], nums[idx]
}
