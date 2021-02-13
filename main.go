package main

import "fmt"

func main() {
	var value, mask, result int

	//(compare each position)
	//Bitwise AND
	value = 63
	mask = 84
	result = value & mask //On: 1, Off: 0
	fmt.Printf("%08b & %08b : %08b\n", value, mask, result)

	//Bitwise OR
	value = 106
	mask = 45
	result = value | mask
	fmt.Printf("%08b | %08b : %08b\n", value, mask, result)

	//Bitwise XOR
	value = 34
	mask = 22
	result = value ^ mask
	fmt.Printf("%08b ^ %08b : %08b\n", value, mask, result)
}
