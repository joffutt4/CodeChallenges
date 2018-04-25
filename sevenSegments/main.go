package main

import (
	"fmt"
	"strings"
)

var segments = map[[3]string]int{
	[3]string{" _ ", "| |", "|_|"}: 0,
	[3]string{"   ", "  |", "  |"}: 1,
	[3]string{" _ ", " _|", "|_ "}: 2,
	[3]string{" _ ", " _|", " _|"}: 3,
	[3]string{"   ", "|_|", "  |"}: 4,
	[3]string{" _ ", "|_ ", " _|"}: 5,
	[3]string{" _ ", "|_ ", "|_|"}: 6,
	[3]string{" _ ", "  |", "  |"}: 7,
	[3]string{" _ ", "|_|", "|_|"}: 8,
	[3]string{" _ ", "|_|", " _|"}: 9,
}

func main() {
	for _, val := range inputs {
		decipher(val)
	}
}

func decipher(segment string) {
	lines := strings.Split(segment, "\n")
	for i := 0; i < len(lines[0]); i += 3 {
		num := [3]string{lines[0][i : i+3], lines[1][i : i+3], lines[2][i : i+3]}
		fmt.Print(segments[num])
	}
	fmt.Println()
}

var inputs = []string{
	`    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|`,
	`    _  _  _  _  _  _  _  _ 
|_| _| _||_|| ||_ |_| _||_ 
  | _| _||_||_| _||_||_  _|`,
	` _  _  _  _  _  _  _  _  _ 
|_  _||_ |_| _|  ||_ | ||_|
 _||_ |_||_| _|  ||_||_||_|`,
	` _  _        _  _  _  _  _ 
|_||_ |_|  || ||_ |_ |_| _|
 _| _|  |  ||_| _| _| _||_ `,
}
