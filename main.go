package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	dichotomies := getDichotomiesFromTSV(input)
}

func getDichotomiesFromTSV(input string) *Dichotomy {
	in, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer in.Close()
	ds := bufio.NewScanner(in)
	for ds.Scan() {
		fmt.Println(ds.Text())
	}
	return &Dichotomy{}
}
