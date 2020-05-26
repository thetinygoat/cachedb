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
	"time"
)

// Set -> struct to hold set data
type Set struct {
	Data    string
	TTL     time.Duration
	AddedAt time.Time
	mutex   sync.Mutex
}

// SMap -> map to hold refs to sets
var SMap = make(map[string]*Set)

func (s *Set) set(value string, ttl int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Data = value
	s.TTL = time.Duration(ttl) * time.Second
	s.AddedAt = time.Now()
}

func (s *Set) get() (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	now := time.Now()
	expiration := s.AddedAt.Add(s.TTL)
	if now.Sub(expiration) > 0 {
		return nilString, errors.New(keyExpired)
	}
	return s.Data, nil
}
