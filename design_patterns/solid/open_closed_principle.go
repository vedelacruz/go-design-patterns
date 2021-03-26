package solid

/*
	makes the interface open for extensibility but closed for modification
*/

type Color int

const (
	Orange Color = iota
	Blue
	Green
	Yellow
	Purple
	Red
)

type Product struct {
	Name  string
	Price float64
	Color Color
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	Color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.Color == c.Color
}

type PriceRangeSpecification struct {
	Price float64
}

func (pr PriceRangeSpecification) IsSatisfied(p *Product) bool {
	return p.Price <= pr.Price
}

func Filter(p []Product, s Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range p {
		if s.IsSatisfied(&v) {
			result = append(result, &p[i])
		}
	}

	return result
}
