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
	d := getDichotomiesFromTSV(input)
	ds, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ds))
}

func getDichotomiesFromTSV(input string) *dichotomies.Dichotomies {
	in, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer in.Close()
	ds := bufio.NewScanner(in)
	answers := make([]*dichotomies.Dichotomy, 0)
	ds.Scan()
	for ds.Scan() {
		d, err := dichotomies.ParseDichotomyFromTSV(ds.Text())
		if err != nil {
			panic(err)
		}
		answers = append(answers, d)
	}
	return dichotomies.NewDichotomies(answers)
}
