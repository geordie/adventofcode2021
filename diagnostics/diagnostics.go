package diagnostics

type Counter int
type FilterType int

const (
	O2 FilterType = iota
	CO2
)

type Diagnostic [12]int

type Diagnostics []Diagnostic

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

func (d *Diagnostic) Epsilon(tot Counter) Diagnostic {
	var dEpsilon Diagnostic
	for i, elem := range d {
		if elem < int(tot)/2 {
			dEpsilon[i] = 1
		}
	}
	return dEpsilon
}

func (d *Diagnostic) Int() int {
	result := 0
	for _, elem := range d {
		result = result << 1
		result = result | elem
	}
	return result
}

func (dArr Diagnostics) Filter(idx int, ft FilterType) Diagnostics {

	// If the array has one element or less, we're done
	if len(dArr) <= 1 {
		return dArr
	}

	// Find the value that occurs most often
	ones, zeros := 0, 0
	var dArrO2 Diagnostics

	for _, elem := range dArr {
		if int(elem[idx]) == 1 {
			ones++
		} else if int(elem[idx]) == 0 {
			zeros++
		}
	}

	// Set a boolean to T if one is more comment, F otherwise
	bMoreOnes := ones > zeros

	// Return the subset of values that qualify
	for _, elem := range dArr {

		if bMoreOnes && elem[idx] == 1 {
			dArrO2 = append(dArrO2, elem)
		} else if !bMoreOnes && elem[idx] == 0 {
			dArrO2 = append(dArrO2, elem)
		}
	}

	// Recurse
	return dArrO2.Filter(idx+1, ft)
}
