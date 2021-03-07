package design_patterns

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0
type Journal struct {
	Entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.Entries = append(j.Entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...remove entry
}

// separation of concerns (i.e. instead of adding these methods to Journal receiver):
var LineSeparator = "\n"

// 1st way create function accepts type
func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.Entries, LineSeparator)), 0644)
}

// 2nd way create a method on receiver
type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.Entries, p.LineSeparator)), 0644)
}