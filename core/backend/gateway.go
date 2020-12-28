/*
 * Copyright (C) 2020 The ocluso Authors
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

package backend

import (
	"log"
	"net/http"
)

type Gateway struct {
	config          *Config
	moduleIndex     *ModuleIndex
	moduleInstances map[string]Module
}

func NewGateway(config *Config, modules *ModuleIndex) (*Gateway, error) {
	result := Gateway{
		config:          config,
		moduleIndex:     modules,
		moduleInstances: make(map[string]Module, 0),
	}

	for _, moduleName := range modules.LoadedModules() {
		indexEntry, err := modules.EntryFor(moduleName)
		if err != nil {
			return nil, err
		}

		context := ModuleContext{
			Configuration: config,
			ModuleIndex:   modules,
		}

		moduleInstance, err := indexEntry.ModuleFactory(context)
		if err != nil {
			return nil, err
		}

		result.moduleInstances[moduleName] = moduleInstance
	}

	return &result, nil
}

func (g *Gateway) Run() {
	moduleMux := NewModuleMux(&g.moduleInstances)

	mux := http.NewServeMux()
	mux.Handle("/", http.RedirectHandler("/app/", 301))
	mux.Handle("/app/", http.FileServer(http.Dir("frontend")))
	mux.Handle("/api/", http.StripPrefix("/api", moduleMux))
	mux.Handle("/api", http.StripPrefix("/api", moduleMux))

	log.Fatal(http.ListenAndServe(":8080", mux)) //TODO: Make port configurable
}
