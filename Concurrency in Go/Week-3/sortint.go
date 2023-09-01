package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

/*
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/

func sortInt(nums []int, ch chan<- []int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("[]int goroutine #%d will sort: %v\n", id, nums)
	sort.Ints(nums)
	ch <- nums
	fmt.Printf("Sorted []int on goroutine #%d: %v\n", id, nums)
}

func main() {
	var wg sync.WaitGroup
	var intSlice []int

	fmt.Printf("Enter a series of integer: ")
	reader := bufio.NewReader(os.Stdin)
	intStr, _ := reader.ReadString('\n')
	for _, numStr := range strings.Fields(intStr) {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			intSlice = append(intSlice, num)
		}
	}

	n := len(intSlice)
	ch := make(chan []int, 4)
	firstCut := n / 4
	secondCut := firstCut * 2
	thirdCut := secondCut + firstCut

	wg.Add(4)
	go sortInt(intSlice[:firstCut], ch, &wg, 1)
	go sortInt(intSlice[firstCut:secondCut], ch, &wg, 2)
	go sortInt(intSlice[secondCut:thirdCut], ch, &wg, 3)
	go sortInt(intSlice[thirdCut:], ch, &wg, 4)
	wg.Wait()

	first, second, third, fourth := <-ch, <-ch, <-ch, <-ch

	fmt.Printf("Final sorted []int: %v\n", merge(merge(first, second), merge(third, fourth)))
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
