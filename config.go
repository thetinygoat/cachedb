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
	"fmt"
	"os"
	"path/filepath"
)

const (
	CONFIG_DIR  = ".config/cachedb"
	CONFIG_FILE = "config.json"
)

type Config struct {
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func initConfigFile() {
	dirPath := filepath.Join(os.Getenv("HOME"), CONFIG_DIR)
	configFilePath := filepath.Join(dirPath, CONFIG_FILE)
	defaultConfig := Config{Port: 9898, User: "", Password: ""}
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
	fmt.Println(config)
}
