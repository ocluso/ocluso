/*
 * Copyright (C) 2021 The ocluso Authors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package core

import (
	"encoding/json"
	"io/ioutil"
)

// Config contains the configuration needed to run the ocluso backend
type Config struct {
	// The address that the HTTP server shall listen on
	HTTPListenAddress string `json:"httpListenAddress"`

	// The data source name for connecting to PostgreSQL
	PostgresDSN string `json:"postgresDSN"`
}

// LoadConfig loads a Config object from a JSON file
func LoadConfig(filename string) (*Config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := Config{}

	err = json.Unmarshal(bytes, &config) //TODO: Write unit test
	if err != nil {
		return nil, err
	}

	return &config, nil
}
