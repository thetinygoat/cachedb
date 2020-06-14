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
	value string
	next  *element
	prev  *element
}

// New instantiates a new list
func New() *List {
	list := &List{}
	return list
}

// Append appends a value to the list
func (list *List) Append(values ...string) {
	for _, value := range values {
		newElement := &element{value: value, prev: list.tail}
		if list.size == 0 {
			list.head = newElement
			list.tail = newElement
		} else {
			list.tail.next = newElement
			list.tail = newElement
		}
		list.size++
	}
}

// Prepend prepends a value to the list
func (list *List) Prepend(values ...string) {
	for v := len(values) - 1; v >= 0; v-- {
		newElement := &element{value: values[v], next: list.head}
		if list.size == 0 {
			list.head = newElement
			list.tail = newElement
		} else {
			list.head.prev = newElement
			list.head = newElement
		}
		list.size++
	}
}

// GetFirst returns first element from the list
func (list *List) GetFirst() (string, bool) {
	if list.size == 0 {
		return "", false
	}
	return list.head.value, true
}

// GetLast returns first element from the list
func (list *List) GetLast() (string, bool) {
	if list.size == 0 {
		return "", false
	}
	return list.tail.value, true
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
