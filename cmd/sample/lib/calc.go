package lib

// Calc struct
type Calc struct {
	V1, V2 int
}

// Add method
func (p *Calc) Add() int {
	return p.V1 + p.V2
}

// Sub method
func (p *Calc) Sub() int {
	return p.V1 - p.V2
}
