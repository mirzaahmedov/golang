package main

import "fmt"

func printList(list []string) {

	fmt.Println("\n\n|index\t|value\n|\t|")
	for i , elem := range list {
		fmt.Printf("|%d.\t|%s\n", i , elem)
	}

}

func main() {

	var input []string = []string{ "a", "a", "a", "b", "c", "d", "a", "b", "g" }
	var output []string;

	outer:
	 for i, elem := range input {
		for _, elem2 := range input[i+1:] {
			for _, elem3 := range output {
				if elem == elem3 {
					continue outer	
				}
			}
			if elem == elem2 {
				output = append(output, elem2)
			}
		}
	}
	printList(output)
}