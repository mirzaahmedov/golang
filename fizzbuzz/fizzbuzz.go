package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getFizzBuzz(length int64) []string {
	var numbers []string
	var count int = 1 
	for len(numbers) <= int(length) {
		if count % 3 == 0 && count % 5 == 0 {
			numbers = append(numbers, "FizzBuzz")
		} else if count % 3 == 0 {
			numbers = append(numbers, "Fizz")
		} else if count % 5 == 0 {
			numbers = append(numbers, "Buzz")
		} else {
			numbers = append(numbers, strconv.Itoa(count))
		}
		count++
	}

	return numbers
}

func printList(list []string) {

	fmt.Println("\n\n|index\t|value\t|\n|\t|\t|")
	for i , elem := range list {
		fmt.Printf("|%d.\t|%s\t|\n", i , elem)
	}

}

func main() {

	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Type any nunber: ")
	scanner.Scan()
	input , _ := strconv.ParseInt(scanner.Text(), 10, 64) 
	var output []string = getFizzBuzz(input)
	printList(output)

}