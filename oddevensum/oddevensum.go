package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getSums(length int64) (odd int, even int) {

	for i := 1; i <= int(length); i++ {
		if i % 2 == 0 {
			even += i
		} else {
			odd += i
		}
	}
	return
}

func main() {

	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Type any nunber: ")
	scanner.Scan()
	input , _ := strconv.ParseInt(scanner.Text(), 10, 64) 
	var odd, even int = getSums(input)
	fmt.Println("sum of all odd numbers is: ", odd)
	fmt.Println("sum of all even numbers is: ", even)

}