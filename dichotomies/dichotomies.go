package dichotomies

import (
	"fmt"
	"strings"
)

//Dichotomy holds the base dichotomy fields, parsed from tsv or database
type Dichotomy struct {
	Order       string `json:"order"`
	Chaos       string `json:"chaos"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

//Dichotomies is the outer object for storing the list of dichotomoies
type Dichotomies struct {
	Dichotomies []*Dichotomy `json:"dichotomies"`
}

//NewDichotomies wraps a collection of dichotomies
func NewDichotomies(d []*Dichotomy) *Dichotomies {
	return &Dichotomies{
		Dichotomies: d,
	}
}

//ParseDichotomyFromTSV takes a raw TSV line and returns a pointer to a Dichotomy struct
func ParseDichotomyFromTSV(line string) (*Dichotomy, error) {
	data := strings.Split(line, "\t")
	if len(data) != 4 {
		return nil, fmt.Errorf("input line parsed into %v fields, but 4 fields are required", len(data))
	}
	return &Dichotomy{
		Order:       data[0],
		Chaos:       data[1],
		Author:      data[2],
		Description: data[3],
	}, nil
}
