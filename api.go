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
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	KEY   = "key"
	VALUE = "value"
	TTL   = "ttl"
	NAME  = "name"
	START = "start"
	STOP  = "stop"
)

type ChannelResponse struct {
	Data  interface{}
	Error error
}

func getSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName(KEY)
	if _, exists := SMap[key]; !exists {
		w.Write([]byte(nilString))
	} else {
		responseChannel := make(chan ChannelResponse)
		go SMap[key].get(responseChannel)
		data := <-responseChannel
		if data.Error != nil {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
			delete(SMap, key)
		} else {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		}
	}
}

func setSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName(KEY)
	if _, exists := SMap[key]; !exists {
		SMap[key] = &Set{}
	}

	responseChannel := make(chan ChannelResponse)
	value := r.URL.Query().Get(VALUE)
	ttlRaw := r.URL.Query().Get(TTL)
	ttl, _ := strconv.Atoi(ttlRaw)
	go SMap[key].set(value, ttl, responseChannel)
	data := <-responseChannel
	if data.Error != nil {
		w.Write([]byte(fmt.Sprintf("%v", data.Data)))
	} else {
		w.Write([]byte(fmt.Sprintf("%v", data.Data)))
	}
}

func delSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName(KEY)
	delete(SMap, key)
	w.Write([]byte(ok))
}

func rpushList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName(NAME)
	if _, exists := LMap[name]; !exists {
		LMap[name] = &List{Name: name}
	}
	responseChannel := make(chan ChannelResponse)
	value := r.URL.Query().Get(VALUE)
	go LMap[name].rpush(value, responseChannel)
	data := <-responseChannel
	if data.Error != nil {
		w.Write([]byte(fmt.Sprintf("%v", data.Data)))
	} else {
		w.Write([]byte(fmt.Sprintf("%v", data.Data)))
	}
}

func lpushList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName(NAME)
	if _, exists := LMap[name]; !exists {
		LMap[name] = &List{Name: name}
	}
	responseChannel := make(chan ChannelResponse)
	value := r.URL.Query().Get(VALUE)
	go LMap[name].lpush(value, responseChannel)
	data := <-responseChannel
	if data.Error != nil {
		w.Write([]byte(fmt.Sprintf("%v", data.Data)))
	} else {
		w.Write([]byte(fmt.Sprintf("%v", data.Data)))
	}
}

func rpopList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName(NAME)
	if _, ok := LMap[name]; !ok {
		w.Write([]byte(nilString))
	} else {
		responseChannel := make(chan ChannelResponse)
		go LMap[name].rpop(responseChannel)
		data := <-responseChannel
		if data.Error != nil {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		} else {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		}
	}
}

func lpopList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName(NAME)
	if _, exists := LMap[name]; !exists {
		w.Write([]byte(nilString))
	} else {
		responseChannel := make(chan ChannelResponse)
		go LMap[name].lpop(responseChannel)
		data := <-responseChannel
		if data.Error != nil {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		} else {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		}
	}
}

func lrangeList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName(NAME)
	if _, exists := LMap[name]; !exists {
		w.Write([]byte(nilString))
	} else {
		startRaw := r.URL.Query().Get(START)
		stopRaw := r.URL.Query().Get(STOP)
		start, _ := strconv.Atoi(startRaw)
		stop, _ := strconv.Atoi(stopRaw)
		responseChannel := make(chan ChannelResponse)
		go LMap[name].lrange(start, stop, responseChannel)
		data := <-responseChannel
		if data.Error != nil {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		} else {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		}
	}
}
