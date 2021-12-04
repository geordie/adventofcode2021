package diagnostics

import (
	"fmt"
	"strconv"
)

type Counter int

type Diagnostic [12]int

func ParseDiagnostic(s string) Diagnostic {
	var d Diagnostic
	for i, c := range s {
		if string(c) == "1" {
			d[i]++
		}
	}
	return d
}

func (mainD *Diagnostic) Add(d Diagnostic) {
	for i, elem := range d {
		mainD[i] += elem
	}
}

func (d *Diagnostic) Gamma(tot Counter) Diagnostic {
	var dGamma Diagnostic
	for i, elem := range d {
		if elem > int(tot)/2 {
			dGamma[i] = 1
		}
	}
	return dGamma
}

func (d *Diagnostic) Int() int {
	result := 0
	for _, elem := range d {
		result = result << 1
		result = result | elem
		fmt.Println(elem, strconv.FormatInt(int64(result), 2))
	}
	return result
}