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
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Router -> holds router config
type Router struct {
	router *httprouter.Router
}

func rpushHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if _, ok := LMap[name]; !ok {
		LMap[name] = &List{Name: name}
	}
	data := r.URL.Query().Get("data")
	err := LMap[name].rpush(data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, ok)
	}
}

func rpopHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if _, ok := LMap[name]; !ok {
		LMap[name] = &List{Name: name}
	}
	data, err := LMap[name].rpop()
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, data)
	}
}

func lpushHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if _, ok := LMap[name]; !ok {
		LMap[name] = &List{Name: name}
	}
	data := r.URL.Query().Get("data")
	err := LMap[name].lpush(data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, ok)
	}
}

func lpopHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if _, ok := LMap[name]; !ok {
		LMap[name] = &List{Name: name}
	}
	data, err := LMap[name].lpop()
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, data)
	}
}

func lrangeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if _, ok := LMap[name]; !ok {
		LMap[name] = &List{Name: name}
	}
	startRaw := r.URL.Query().Get("start")
	stopRaw := r.URL.Query().Get("stop")
	start, err := strconv.Atoi(startRaw)
	stop, err := strconv.Atoi(stopRaw)
	data, err := LMap[name].lrange(start, stop)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, data)
	}
}

func setHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	if _, ok := SMap[key]; !ok {
		SMap[key] = &Set{}
	}
	value := r.URL.Query().Get("value")
	ttlRaw := r.URL.Query().Get("ttl")
	ttl, _ := strconv.Atoi(ttlRaw)
	SMap[key].set(value, ttl)
	fmt.Fprint(w, ok)
}
func getHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	if _, ok := SMap[key]; !ok {
		fmt.Fprint(w, nilString)
		return
	}
	data, err := SMap[key].get()
	if err != nil {
		fmt.Fprint(w, err.Error())
		delete(SMap, key)
	} else {
		fmt.Fprint(w, data)
	}
}

func delHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	delete(SMap, key)
	fmt.Fprint(w, ok)

}

func (R *Router) initializeRouter() {
	R.router = httprouter.New()
	R.router.POST("/lists/rpush/:name", rpushHandler)
	R.router.POST("/lists/rpop/:name", rpopHandler)
	R.router.POST("/lists/lpush/:name", lpushHandler)
	R.router.POST("/lists/lpop/:name", lpopHandler)
	R.router.GET("/lists/lrange/:name", lrangeHandler)
	R.router.POST("/sets/set/:key", setHandler)
	R.router.GET("/sets/get/:key", getHandler)
	R.router.POST("/sets/del/:key", delHandler)
	log.Fatal(http.ListenAndServe(":9898", R.router))
}
