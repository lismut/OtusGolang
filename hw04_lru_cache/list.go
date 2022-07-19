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
	Value interface {}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	First  *ListItem
	Last   *ListItem
	Length int
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *ListItem {
	return l.First
}

func (l *list) Back() *ListItem {
	return l.Last
}

func (l *list) PushFront(v interface{}) *ListItem {
	NewFront := ListItem{v, nil, nil}
	if l.Front() != nil {
		l.Front().Prev = &NewFront
		NewFront.Next = l.Front()
	} else {
		l.Last = &NewFront
	}
	l.First = &NewFront
	l.Length++
	return l.Front()
}

func (l *list) PushBack(v interface{}) *ListItem {
	NewBack := ListItem{v, nil, nil}
	if l.Back() != nil {
		l.Back().Next = &NewBack
		NewBack.Prev = l.Back()
	} else {
		l.First = &NewBack
	}
	l.Last = &NewBack
	l.Length++
	return l.Back()
}

func (l *list) Remove(i *ListItem) {
	if i.Value != nil {
		switch {
		case i.Prev == nil && i.Next == nil:
			l.First = nil
			l.Last = nil
		case i.Prev == nil:
			i.Next.Prev = nil
			l.First = i.Next
		case i.Next == nil:
			i.Prev.Next = nil
			l.Last = i.Prev
		default:
			i.Next.Prev = i.Prev
			i.Prev.Next = i.Next
		}

		l.Length--
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove((i))
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
