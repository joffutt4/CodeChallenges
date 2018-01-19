package main

import (
	"fmt"
)

func main() {
	taps := [2]int{1, 2}
	lfsr(taps[:], "XOR", "001", 7)
	taps = [2]int{0, 2}
	lfsr(taps[:], "XNOR", "001", 7)
	taps4 := [4]int{1, 2, 3, 7}
	lfsr(taps4[:], "XOR", "00000001", 16)
	taps4 = [4]int{1, 5, 6, 31}
	lfsr(taps4[:], "XOR", "00000000000000000000000000000001", 16)
}

func lfsr(taps []int, op, start string, iter int) {
	bytes := []byte(start)
	var ins byte
	fmt.Println(fmt.Sprintf("%d %s", 0, start))
	for i := 1; i <= iter; i++ {
		ins = 0
		for j := 0; j < len(taps); j++ {
			if op == "XOR" {
				ins ^= bytes[taps[j]]
			} else {
				ins = ^(ins ^ bytes[taps[j]])
			}
		}
		// After performing binary operations on the bytes, they will either be 0 or 1 so we will
		// add 48 to the byte value since the string char byte value for 0 is 48.
		// Doing so ensure that the bytes in the array cleanly translate to the chars for 0 (48 byte) and 1 (49 byte)
		// when this byte array is converted to a string
		ins += 48
		// Remove the last element from the array
		bytes = bytes[:len(bytes)-1]
		// Insert the newest element to the front of the array
		bytes = append([]byte{ins}, bytes...)
		val := string(bytes[:])
		fmt.Println(fmt.Sprintf("%d %s", i, val))
	}
}
