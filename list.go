// Copyright (C) 2020  Sachin Saini

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

import (
	"errors"
	"strings"
)

// ListNode -> Node struct
type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

// List -> List struct
type List struct {
	name string
	head *ListNode
	tail *ListNode
	size int
}

// LMap -> List map to hold refs to all the lists
var LMap = make(map[string]*List)

func (l *List) append(data string, response chan ChannelResponse) {
	newNode := &ListNode{data: data, prev: l.tail}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
	response <- ChannelResponse{Data: ok, Error: nil}
}

func (l *List) prepend(data string, response chan ChannelResponse) {
	newNode := &ListNode{data: data, next: l.head}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.prev = newNode
		l.head = newNode
	}
	l.size++
	response <- ChannelResponse{Data: ok, Error: nil}
}

func (l *List) removelast(response chan ChannelResponse) {

	if l.size == 0 {
		response <- ChannelResponse{Data: nilString, Error: errors.New(listEmpty)}
	} else if l.size == 1 {
		data := l.head.data
		l.clear()
		response <- ChannelResponse{Data: data, Error: nil}
	} else {
		node := l.tail
		data := node.data
		if l.tail.prev != nil {
			l.tail = l.tail.prev
		}
		node = nil
		l.size--
		response <- ChannelResponse{Data: data, Error: nil}
	}
}

func (l *List) removefirst(response chan ChannelResponse) {
	if l.size == 0 {
		response <- ChannelResponse{Data: nilString, Error: errors.New(listEmpty)}
	} else if l.size == 1 {
		data := l.head.data
		l.clear()
		response <- ChannelResponse{Data: data, Error: nil}
	} else {
		node := l.head
		data := node.data
		if l.head.next != nil {
			l.head = l.head.next
		}
		node = nil
		l.size--
		response <- ChannelResponse{Data: data, Error: nil}
	}
}

func (l *List) values(response chan ChannelResponse) {
	if l.size == 0 {
		response <- ChannelResponse{Data: nilString, Error: errors.New(listEmpty)}
	} else {
		startRef := l.head
		var b strings.Builder
		for startRef != nil {
			b.WriteString(startRef.data)
			b.WriteString(" ")
			startRef = startRef.next
		}
		response <- ChannelResponse{Data: b.String(), Error: nil}
	}
}

func (l *List) clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}
