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

func appendHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	append(w, r, ps)
}

func prependHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	prepend(w, r, ps)
}

func removelastHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	removelast(w, r, ps)
}

func removefirstHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	removefirst(w, r, ps)
}

func valuesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	values(w, r, ps)
}

func (R *Router) initializeRouter() {
	R.router = httprouter.New()
	R.router.POST("/lists/append/:name", appendHandler)
	R.router.POST("/lists/prepend/:name", prependHandler)
	R.router.POST("/lists/removelast/:name", removelastHandler)
	R.router.POST("/lists/removefirst/:name", removefirstHandler)
	R.router.GET("/lists/values/:name", valuesHandler)
	R.router.POST("/sets/set/:key", setHandler)
	R.router.GET("/sets/get/:key", getHandler)
	R.router.POST("/sets/del/:key", delHandler)
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv(CachedbPort)),
		Handler: Logger(Auth(R.router)),
	}
	fmt.Printf("Cachedb server running on port: %s\n", os.Getenv(CachedbPort))
	log.Fatal(s.ListenAndServe())
}
