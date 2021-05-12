package main

import "fmt"

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func printList(list []string) {

	fmt.Println("\n\n|index\t|value\n|\t|")
	for i , elem := range list {
		fmt.Printf("|%d.\t|%s\n", i , elem)
	}

}

func main() {

	var input []string = []string{ "dog", "cow", "tap", "god", "pat", "abcd", "dcba", "lls", "s", "sssll" }
	var output []string;

	for _, elem := range input {
		for _, elem2 := range input {
			var concat string = elem + elem2
			if len(concat) % 2 == 0 {
				if concat[:len(concat)/2] == Reverse(concat[len(concat)/2:]) {
					output = append(output, concat)
			   }
			} else {
				if concat[:len(concat)/2] == Reverse(concat[len(concat)/2+1:]) {
					output = append(output, concat)
				}
			}
		}
	}
	printList(output)
}