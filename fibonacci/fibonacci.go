package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getFibonacci(length int64) []int {
	var numbers []int = []int{0, 1}
	for len(numbers) <= int(length) {
		numbers = append(numbers, numbers[len(numbers) - 2] + numbers[len(numbers) - 1])
	}

	return numbers
}

func printList(list []int) {

	fmt.Println("\n\n|index\t|value\t|\n|\t|\t|")
	for i , elem := range list {
		fmt.Printf("|%d.\t|%d\t|\n", i , elem)
	}

}

func main() {

	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Type any nunber: ")
	scanner.Scan()
	input , _ := strconv.ParseInt(scanner.Text(), 10, 64) 
	var output []int = getFibonacci(input)
	printList(output)

}