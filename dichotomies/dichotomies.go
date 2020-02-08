package dichotomies

import (
	"fmt"
	"strings"
)

//Dichotomy holds the base dichotomy fields, parsed from tsv or database
type Dichotomy struct {
	Order       string `json:"order"`
	Chaos       string `json:"chaos"`
	Originator  string `json:"originator"`
	Description string `json:"description"`
}

//ParseDichotomyFromTSV takes a raw TSV line and returns a pointer to a Dichotomy struct
func ParseDichotomyFromTSV(line string) (*Dichotomy, error) {
	data := strings.Split(line, "	")
	if len(data) != 4 {
		return nil, fmt.Errorf("input line parsed into %v fields, but 4 fields are required", len(data))
	}
	return &Dichotomy{
		Order:       data[0],
		Chaos:       data[1],
		Originator:  data[2],
		Description: data[3],
	}, nil
}
