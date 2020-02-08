package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/plantimals/dichotomies/dichotomies"
)

func main() {
	input := os.Args[1]
	ds := getDichotomiesFromTSV(input)
	for _, d := range ds {
		fmt.Println(json.Marshal(d))
	}
}

func getDichotomiesFromTSV(input string) []*dichotomies.Dichotomy {
	in, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer in.Close()
	ds := bufio.NewScanner(in)
	answers := make([]*dichotomies.Dichotomy, 0)
	for ds.Scan() {
		d, err := dichotomies.ParseDichotomyFromTSV(ds.Text())
		if err != nil {
			panic(err)
		}
		answers = append(answers, d)
	}
	return answers
}
