// ex 4.3 - Rescreva reverse usando um ponteiro de array no lugar de uma fatia
package main

import "fmt"

func main() {
	arrayExample := [...]int{1998, 1999, 2000, 2001, 2002, 2003}
	fmt.Println(arrayExample)

	reverse(&arrayExample)
	fmt.Println(arrayExample)

}

func reverse(ints *[6]int) {
	for i := 0; i < len(ints)/2; i++ {
		reversedIndex := len(ints) - 1 - i
		ints[i], ints[reversedIndex] = ints[reversedIndex], ints[i]
	}
}
