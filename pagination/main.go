package main

import "fmt"

// Pages contains a slice of strings that will hold the information
type Pages struct {
	Elements []string
	// PageSize specifies the number of elements that a single page will hold
	PageSize int
	// Current specifies which page we currently are on
	Current int
}

// FirstPage returns the elements on the first page and sets the current page to first
func (p *Pages) FirstPage() []string {
	p.Current = 1
	fmt.Println(p.Current)
	s := p.Elements[:p.PageSize]
	return s
}

// LastPage returns the elements on the last page and sets the current page to last
func (p *Pages) LastPage() []string {
	p.Current = len(p.Elements) / p.PageSize
	if len(p.Elements)%p.PageSize != 0 {
		p.Current++
	}
	s := p.Elements[len(p.Elements)-len(p.Elements)%p.PageSize:]
	return s
}

// GetCurrentPageNumber returns the current page that we are on
func (p *Pages) GetCurrentPageNumber() int {
	return p.Current
}

// HasNext checks if there is a page after the one we are currently on
func (p Pages) HasNext() bool {
	totalPages := len(p.Elements) / p.PageSize
	if len(p.Elements)%p.PageSize != 0 {
		totalPages++
	}
	if p.Current == totalPages {
		return false
	}
	return true
}

// HasPrevious checks if there is a page before the one we are currently on
func (p Pages) HasPrevious() bool {
	if p.Current == 1 {
		return false
	}
	return true
}

// Next returns the page after the one we are currently on
func (p *Pages) Next() []string {
	s := make([]string, p.PageSize)
	if p.HasNext() != true {
		return s
	}
	for i := p.Current * p.PageSize; i < p.Current*p.PageSize+p.PageSize; i++ {
		s = append(s, p.Elements[i])
		if i == len(p.Elements)-1 {
			break
		}
	}
	return s
}

// Previous returns the page before the one we are currently on
func (p *Pages) Previous() []string {
	s := make([]string, p.PageSize)
	if p.HasPrevious() != true {
		return s
	}
	for i := p.Current*p.PageSize - p.PageSize*2; i < p.Current*p.PageSize-p.PageSize; i++ {
		s = append(s, p.Elements[i])
		if i == len(p.Elements)+1 {
			break
		}
	}
	return s
}

func main() {
	s := Pages{[]string{"book", "work", "cat", "house", "phone", "case", "cup", "mouse", "watch", "clock", "dog"}, 3, 3}
	fmt.Println(s)
	fmt.Println(s.FirstPage())
	fmt.Println(s)
	fmt.Println(s.LastPage())
	fmt.Println(s)
	fmt.Println(s.GetCurrentPageNumber())
	fmt.Println(s.HasNext())
	fmt.Println(s.HasPrevious())
	fmt.Println(s.Next())
	fmt.Println(s.Previous())
}
