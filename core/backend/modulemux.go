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

import "net/http"

// ModuleMux is an HTTP multiplexer for ocluso Modules
type ModuleMux struct {
	mux *http.ServeMux
}

// NewModuleMux creates a new ModuleMux
//
// The modules argument is a map with the module name as key and the Module
// instance as value.
//
// Each ocluso Module has its module name as its own HTTP subpath.
func NewModuleMux(modules *map[string]Module) *ModuleMux {
	result := ModuleMux{
		mux: http.NewServeMux(),
	}

	for moduleName, module := range *modules {
		prefix := "/" + moduleName + "/"
		result.mux.Handle(prefix, http.StripPrefix(prefix, module))
	}

	return &result
}

func (m *ModuleMux) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	m.mux.ServeHTTP(responseWriter, request)
}
