package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	counter int
	first   *ListItem
	last    *ListItem
}

func (l *list) Len() int {
	return l.counter
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	first := &ListItem{v, l.first, nil}
	if l.counter == 1 { // if list of one element - move first to last
		l.last = l.first
		l.last.Next = nil
		l.last.Prev = first
	} else if l.counter >= 2 { // if list of more than 1 element
		l.first.Prev = first
	}
	l.first = first
	l.counter++
	return first
}

func (l *list) PushBack(v interface{}) *ListItem {
	last := &ListItem{v, nil, l.last}
	if l.first == nil {
		l.first = last
		l.counter++
		return last
	}
	if l.last == nil {
		l.first.Next = last
		last.Prev = l.first
	}
	if l.last != nil {
		l.last.Next = last
	}
	l.last = last
	l.counter++
	return last
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Prev == nil && i.Next == nil: // if it is first element and the only one
		l.first = nil
	case i.Prev == nil && l.counter == 2: // or first of a pair
		l.first = l.last
		l.last = nil
		l.first.Next = nil
		l.first.Prev = nil
	case i.Prev == nil && l.counter > 2: // or first of more than 2
		l.first = i.Next
		l.first.Prev = nil
	case i.Next == nil && l.counter == 2: // if it is last element and last of a pair
		l.first.Next = nil
		l.last = nil
	case i.Next == nil && l.counter > 2: // or last of more than 2 elements
		l.last = i.Prev
		l.last.Next = nil
	default: // if it is not last element and not first
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.counter--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil { // if it is first element
		return
	}
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
