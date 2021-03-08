package design_patterns

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) GetWidth() int {
	return r.Width
}

func (r *Rectangle) SetWidth(width int) {
	r.Width = width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) SetHeight(height int) {
	r.Height = height
}

// Squares dont implement Sized interface because SetWidth or SetHeight each would have to set both height and width
// and violates Liskov substitution principle. Implementers of a type should not break core fundamental behaviors
// you rely on.
type Square struct {
	Size int // width, height
}

func (s *Square) Rectangle() Rectangle {
	return Rectangle{s.Size, s.Size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Printf("Expected an area of %d, but got %d\n", expectedArea, actualArea)
}
