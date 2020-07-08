package lib

// Calc2 struct
type Calc2 struct {
	*Calc
}

// Add2 method
func (calc *Calc2) Add2() int {
	return calc.Add()
}