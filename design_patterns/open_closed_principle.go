package design_patterns

/*
	makes the interface open for extensibility but closed for modification
*/

type Color int

const (
	orange Color = iota
	blue
	green
	yellow
	purple
	red
)

type Product struct {
	name  string
	price float64
	color Color
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return c.color == p.color
}

type Filter struct {
	//
}

func (f Filter) Filter(p []Product, s Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range p {
		if s.IsSatisfied(&v) {
			result = append(result, &p[i])
		}
	}

	return result
}
