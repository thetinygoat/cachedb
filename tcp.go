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
