package design_patterns

type Employee struct {
	FirstName, MiddleName, LastName string
}

func (p *Employee) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.FirstName
	}()

	return out
}
