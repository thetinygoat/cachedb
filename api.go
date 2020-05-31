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

// constants
const (
	Key   = "key"
	Value = "value"
	TTL   = "ttl"
	Name  = "name"
	Start = "start"
	Stop  = "stop"
)

// ChannelResponse struct
type ChannelResponse struct {
	Data  interface{}
	Error error
}

func getSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName(Key)
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
	key := ps.ByName(Key)
	if _, exists := SMap[key]; !exists {
		SMap[key] = &Set{}
	}

	responseChannel := make(chan ChannelResponse)
	value := r.URL.Query().Get(Value)
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
	key := ps.ByName(Key)
	delete(SMap, key)
	w.Write([]byte(ok))
}

func append(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName(Name)
	if _, exists := LMap[name]; !exists {
		LMap[name] = &List{name: name}
	}
	responseChannel := make(chan ChannelResponse)
	value := r.URL.Query().Get(Value)
	go LMap[name].append(value, responseChannel)
	data := <-responseChannel
	if data.Error != nil {
		w.Write([]byte(fmt.Sprintf("%v", data.Data)))
	} else {
		w.Write([]byte(fmt.Sprintf("%v", data.Data)))
	}
}

func prepend(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName(Name)
	if _, exists := LMap[name]; !exists {
		LMap[name] = &List{name: name}
	}
	responseChannel := make(chan ChannelResponse)
	value := r.URL.Query().Get(Value)
	go LMap[name].prepend(value, responseChannel)
	data := <-responseChannel
	if data.Error != nil {
		w.Write([]byte(fmt.Sprintf("%v", data.Data)))
	} else {
		w.Write([]byte(fmt.Sprintf("%v", data.Data)))
	}
}

func removelast(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName(Name)
	if _, ok := LMap[name]; !ok {
		w.Write([]byte(nilString))
	} else {
		responseChannel := make(chan ChannelResponse)
		go LMap[name].removelast(responseChannel)
		data := <-responseChannel
		if data.Error != nil {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		} else {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		}
	}
}

func removefirst(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName(Name)
	if _, exists := LMap[name]; !exists {
		w.Write([]byte(nilString))
	} else {
		responseChannel := make(chan ChannelResponse)
		go LMap[name].removefirst(responseChannel)
		data := <-responseChannel
		if data.Error != nil {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		} else {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		}
	}
}

func values(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName(Name)
	if _, exists := LMap[name]; !exists {
		w.Write([]byte(nilString))
	} else {
		responseChannel := make(chan ChannelResponse)
		go LMap[name].values(responseChannel)
		data := <-responseChannel
		if data.Error != nil {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		} else {
			w.Write([]byte(fmt.Sprintf("%v", data.Data)))
		}
	}
}
