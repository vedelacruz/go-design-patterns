package design_patterns

import "fmt"

// break up interface into separate parts that people need

type Document struct {
}

// do not put Print and Scanner in one interface bc not all printers have scan method
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type OldFashionedPrinter struct {
}

func (ofp *OldFashionedPrinter) Print() {
	fmt.Println("...printing on inkjet")
}

type Photocopier struct {
}

func (p *Photocopier) Print(d Document) {
	fmt.Println("...printing on laser jet")
}

func (p *Photocopier) Scan(d Document) {
	fmt.Println("...scanning your file")
}

// can also compose interface
type MultiFunctionDevice interface {
	Printer
	Scanner
}
