// Copyright (C) 2020 Sachin Saini

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

// List holds the elements, which point to next and previous elements
type List struct {
	head *element
	tail *element
	size int
}

type element struct {
	key   string
	value string
	next  *element
	prev  *element
}

// New instantiates a new list
func New() *List {
	return &List{}
}

// Append appends a value to the list
func (list *List) Append(key string, value string) {
	newElement := &element{key: key, value: value, prev: list.tail}
	if list.size == 0 {
		list.head = newElement
		list.tail = newElement
	} else {
		list.tail.next = newElement
		list.tail = newElement
	}
	list.size++

}

// Prepend prepends a value to the list
func (list *List) Prepend(key string, value string) {
	newElement := &element{key: key, value: value, next: list.head}
	if list.size == 0 {
		list.head = newElement
		list.tail = newElement
	} else {
		list.head.prev = newElement
		list.head = newElement
	}
	list.size++

}

// GetFirst returns first element from the list
func (list *List) GetFirst() []string {
	if list.size == 0 {
		return nil
	}
	return []string{list.head.key, list.head.value}
}

// GetFirstRef returns reference to first element from the list
func (list *List) GetFirstRef() *element {
	if list.size == 0 {
		return nil
	}
	return list.head
}

// GetLast returns first element from the list
func (list *List) GetLast() []string {
	if list.size == 0 {
		return nil
	}
	return []string{list.tail.key, list.tail.value}
}

// GetLastRef returns reference to last element from the list
func (list *List) GetLastRef() *element {
	if list.size == 0 {
		return nil
	}
	return list.tail
}

// RemoveFirst removes an element from the front of the list
func (list *List) RemoveFirst() {
	if list.size == 0 {
		return
	}
	if list.size == 1 {
		list.Clear()
		return
	}
	if list.head.next != nil {
		list.head = list.head.next
	}
	list.head.prev = nil
	list.size--
}

// RemoveLast removes an element from the end of the list
func (list *List) RemoveLast() {
	if list.size == 0 {
		return
	}
	if list.size == 1 {
		list.Clear()
		return
	}
	if list.tail.prev != nil {
		list.tail = list.tail.prev
	}
	list.tail.next = nil
	list.size--
}

// Remove removes node from anyhwere in between
func (list *List) Remove(el *element) {
	if list.size == 0 {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	prev := el.prev
	next := el.next

	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	el = nil
}

// Empty returns if list is empty or not
func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns the size of the list
func (list *List) Size() int {
	return list.size
}

// Clear removes all elements from the list
func (list *List) Clear() {
	list.size = 0
	list.head = nil
	list.tail = nil
}
