package hw_6

import "fmt"

type Item struct {
	val  interface{}
	next *Item
	prev *Item
}

func (i Item) Prev() *Item {
	return i.prev
}

func (i Item) Next() *Item {
	return i.next
}

func (i Item) Value() interface{} {
	return i.val
}

type List struct {
	root Item
	len  int
}

func (l *List) Len() int {
	return l.len
}

func (l *List) First() *Item {
	return l.root.next
}

func (l *List) Last() *Item {
	return l.root.prev
}

func (l *List) PushFront(val interface{}) {
	item := Item{
		val:  val,
		prev: &l.root,
		next: l.root.next,
	}
	if l.root.next != nil {
		l.root.next.prev = &item
	}
	l.root.next = &item
	if l.Len() == 0 {
		item.next = &l.root
		l.root.prev = &item
	}
	l.len++
}

func (l *List) PushBack(val interface{}) {
	item := Item{
		val:  val,
		prev: l.root.prev,
		next: &l.root,
	}
	if l.root.prev != nil {
		l.root.prev.next = &item
	}
	l.root.prev = &item
	if l.Len() == 0 {
		item.prev = &l.root
		l.root.next = &item
	}
	l.len++
}

func (l *List) Remove(i *Item) {
	i.next.prev = i.prev
	i.prev.next = i.next
	l.len--
}

func Check() {
	l := List{}
	l.PushFront("c")
	l.PushFront("b")
	l.PushFront("a")
	fmt.Printf("List length. Expected 3. Got %d\n", l.Len())
	fmt.Printf("First. Expected a. Got %s\n", l.First().Value())
	fmt.Printf("First prev. Expected nill. Got %v\n", l.First().Prev().Value())
	fmt.Printf("First prev next. Expected a. Got %s\n", l.First().Prev().Next().Value())
	fmt.Printf("First next. Expected b. Got %s\n", l.First().Next().Value())
	fmt.Printf("First next next. Expected c. Got %s\n", l.First().Next().Next().Value())
	fmt.Printf("First next next next. Expected nil. Got %v\n", l.First().Next().Next().Next().Value())
	fmt.Printf("Last. Expected c. Got %s\n", l.Last().Value())
	fmt.Printf("Last prev. Expected b. Got %s\n", l.Last().Prev().Value())
	fmt.Printf("Last next. Expected nil. Got %v\n", l.Last().Next().Value())
	l.PushBack("z")
	fmt.Printf("List length. Expected 4. Got %d\n", l.Len())
	fmt.Printf("First. Expected a. Got %s\n", l.First().Value())
	fmt.Printf("First prev. Expected nill. Got %v\n", l.First().Prev().Value())
	fmt.Printf("First prev next. Expected a. Got %s\n", l.First().Prev().Next().Value())
	fmt.Printf("First next. Expected b. Got %s\n", l.First().Next().Value())
	fmt.Printf("First next next. Expected c. Got %s\n", l.First().Next().Next().Value())
	fmt.Printf("First next next next. Expected z. Got %s\n", l.First().Next().Next().Next().Value())
	fmt.Printf("Last. Expected z. Got %s\n", l.Last().Value())
	fmt.Printf("Last prev. Expected c. Got %s\n", l.Last().Prev().Value())
	fmt.Printf("Last next. Expected nil. Got %v\n", l.Last().Next().Value())
	l.Remove(l.First().Next())
	fmt.Printf("List length. Expected 3. Got %d\n", l.Len())
	fmt.Printf("First next. Expected c. Got %s\n", l.First().Next().Value())
	fmt.Printf("First next next. Expected z. Got %s\n", l.First().Next().Next().Value())
	fmt.Printf("First next next next. Expected nil. Got %v\n", l.First().Next().Next().Next().Value())
	fmt.Printf("Last. Expected z. Got %s\n", l.Last().Value())
	fmt.Printf("Last prev. Expected c. Got %s\n", l.Last().Prev().Value())
	fmt.Printf("Last prev prev. Expected a. Got %v\n", l.Last().Prev().Prev().Value())
}
