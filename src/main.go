package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the data to be compressed:")
	srcInput, _ := reader.ReadString('\n')

	// Convert string input to byte slice as required by GetSequences
	src := []byte(srcInput)
	numSequences := GetSequences(src)
	fmt.Printf("Number of sequences: %d\n", numSequences)
}
