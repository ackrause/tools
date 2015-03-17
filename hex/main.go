package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	usage := "hex <filename> <number of bytes>"
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(2)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	nbytes, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error converting to number of bytes: %s", err)
		os.Exit(1)
	}
	out := make([]byte, nbytes)

	_, err = file.Read(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s", err)
		os.Exit(1)
	}

	fmt.Printf("%x\n", out)
}
