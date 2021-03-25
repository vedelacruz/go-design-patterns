package design_patterns

import "fmt"

// high level modules should not depend on low level modules
// both should depend on abstractions (i.e. in GO it is interfaces)
// protects when everything breaks down

type Relationship int

const (
	Parent Relationship = iota
	Child
)

type Person struct {
	Name string
}

type Info struct {
	From         *Person
	Relationship Relationship
	To           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// low-level module (closer to hardware, data storage, system level)
type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.Relationship == Parent && v.From.Name == name {
			result = append(result, r.relations[i].To)
		}
	}

	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{
		parent,
		Parent,
		child,
	})

	r.relations = append(r.relations, Info{
		child,
		Child,
		parent,
	})
}

// high-level module (business logic)
type Research struct {
	Browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.Browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.Name)
	}
}
