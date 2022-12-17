package graphics

type Programmable interface {
	Program() Program
}

type programmable struct {
	program Program
}

func (p *programmable) Program() Program {
	return p.program
}
