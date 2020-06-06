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
	"encoding/json"
	"os"
	"path/filepath"
)

// constants
const (
	ConfigDir       = ".config/cachedb"
	ConfigFile      = "config.json"
	CachedbPort     = "CACHEDB_PORT"
	CachedbUser     = "CACHEDB_USER"
	CachedbPassword = "CACHEDB_PASSWORD"
)

// Config struct
type Config struct {
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func initConfigFile() {
	dirPath := filepath.Join(os.Getenv("HOME"), ConfigDir)
	configFilePath := filepath.Join(dirPath, ConfigFile)
	defaultConfig := Config{Port: "9898", User: "cachedb", Password: "cachedb"}
	if _, err := os.Stat(configFilePath); err != nil {
		os.Mkdir(dirPath, 0755)
		f, err := os.Create(configFilePath)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		json.NewEncoder(f).Encode(&defaultConfig)
	}
	f, err := os.Open(configFilePath)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	var config Config
	json.NewDecoder(f).Decode(&config)
	os.Setenv(CachedbPort, config.Port)
	os.Setenv(CachedbUser, config.User)
	os.Setenv(CachedbPassword, config.Password)
}
