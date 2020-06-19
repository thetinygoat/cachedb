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

// Port is the port cachedb runs on
var Port string

// User is the cachedb user
var User string

// Password is the cachedb password
var Password string

// EvictionPolicy is the policy cachedb uses to evict old data
var EvictionPolicy string

// OperationMode is the mode cachedb running in, eg: normal or clustered
var OperationMode string

// MaxMemory is the memory cachedb is allowed to use
var MaxMemory uint64

// ConfigFile provides name for config file
var ConfigFile = "cachedb.conf"

// cache status codes and messages
const (
	StatusNotFound        = 101
	StatusExpired         = 102
	StatusOk              = 103
	StatusMemoryOverload  = 104
	StatusUnexpected      = 105
	MessageNotFound       = "<nil>"
	MessageExpired        = "<exp>"
	MessageOk             = "<ok>"
	MessageMemoryOverload = "<memovld>"
	MessageUnexpected     = "an unexpected error occured"
)

// response messages for handlers
const (
	MessageMalformed = "<malformed>"
)
