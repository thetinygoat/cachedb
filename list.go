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
	"sync"
)

// ListNode -> Node struct
type ListNode struct {
	Data string
	Next *ListNode
	Prev *ListNode
}

// List -> List struct
type List struct {
	Name  string
	Head  *ListNode
	Tail  *ListNode
	Size  int
	mutex sync.Mutex
}

// LMap -> List map to hold refs to all the lists
var LMap = make(map[string]*List)

func (l *List) rpush(data string, response chan ChannelResponse) {
	if l.Head == nil {
		l.Head = &ListNode{Data: data, Next: nil, Prev: nil}
		l.Tail = l.Head
	} else {
		l.Tail.Next = &ListNode{Data: data, Next: nil, Prev: l.Tail}
		l.Tail = l.Tail.Next
	}
	l.Size++
	response <- ChannelResponse{Data: ok, Error: nil}
}

func (l *List) lpush(data string, response chan ChannelResponse) {
	if l.Head == nil {
		l.Head = &ListNode{Data: data, Next: nil, Prev: nil}
		l.Tail = l.Head
	} else {
		node := &ListNode{Data: data, Next: l.Head, Prev: nil}
		l.Head = node
	}
	l.Size++
	response <- ChannelResponse{Data: ok, Error: nil}
}

func (l *List) rpop(response chan ChannelResponse) {
	if l.Head == nil {
		response <- ChannelResponse{Data: nilString, Error: errors.New(listEmpty)}
	} else {
		rdata := l.Tail.Data
		l.Tail = l.Tail.Prev
		if l.Tail != nil {
			l.Tail.Next = nil
		}
		if l.Size == 1 {
			l.Head = nil
			l.Tail = nil
		}
		l.Size--
		response <- ChannelResponse{Data: rdata, Error: nil}
	}
}

func (l *List) lpop(response chan ChannelResponse) {
	if l.Head == nil {
		response <- ChannelResponse{Data: nilString, Error: errors.New(listEmpty)}
	} else {
		rdata := l.Head.Data
		l.Head = l.Head.Next
		if l.Head != nil {
			l.Head.Prev = nil
		}
		if l.Size == 1 {
			l.Head = nil
			l.Tail = nil
		}
		l.Size--
		response <- ChannelResponse{Data: rdata, Error: nil}
	}
}

func (l *List) lrange(start int, stop int, response chan ChannelResponse) {
	if start < 0 {
		response <- ChannelResponse{Data: []string{}, Error: errors.New(invalidRange)}
	} else {
		if stop < 0 {
			stop = l.Size
		}
		if stop > l.Size {
			stop = l.Size
		}
		idx := 0
		startRef := l.Head
		for idx != start {
			startRef = startRef.Next
			idx++
		}
		var values []string
		for start < stop {
			values = append(values, startRef.Data)
			startRef = startRef.Next
			start++
		}
		response <- ChannelResponse{Data: values, Error: nil}
	}
}
