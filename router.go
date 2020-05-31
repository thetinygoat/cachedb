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
	"os"

	"github.com/julienschmidt/httprouter"
)

// Router -> holds router config
type Router struct {
	router *httprouter.Router
}

func setHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setSet(w, r, ps)
}

func getHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	getSet(w, r, ps)
}

func delHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	delSet(w, r, ps)
}

func rpushHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rpushList(w, r, ps)
}

func rpopHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rpopList(w, r, ps)
}

func lpushHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	lpushList(w, r, ps)
}

func lpopHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	lpopList(w, r, ps)
}

func lrangeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	lrangeList(w, r, ps)
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
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv(CachedbPort)),
		Handler: Logger(R.router),
	}
	fmt.Printf("Cachedb server running on port: %s\n", os.Getenv(CachedbPort))
	log.Fatal(s.ListenAndServe())
}
