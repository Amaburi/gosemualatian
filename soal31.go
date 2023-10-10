package main

import (
	"fmt"
)

func main() {
	var char1, char2 byte
	var areEqual bool
	fmt.Scanf("%c %c", &char1, &char2)

	char1 = fmt.Sprintf("%c", char1)[0]
	char2 = fmt.Sprintf("%c", char2)[0]

	areEqual = char1|32 == char2|32

	fmt.Println(areEqual)
}
