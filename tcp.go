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

import (
	"bufio"
	"io"
	"net"
	"strings"
)

func connect(handler *Handler) {
	ln, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn, handler)
	}
}

func handle(conn net.Conn, handler *Handler) {
	for {
		buf, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		buf = buf[:len(buf)-1]
		parsed := requestParser(string(buf))
		var res string
		if parsed[0] == "get" {
			res = handler.handleGet(parsed[1:])
		} else if parsed[0] == "set" {
			res = handler.handleSet(parsed[1:])
		}

		conn.Write([]byte(res + "\n"))
	}
}

func requestParser(data string) []string {
	return strings.Split(data, "%%")
}
