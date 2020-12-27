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

type Module interface {
	http.Handler
}

type ModuleContext struct {
	Configuration *Config
	ModuleIndex   *ModuleIndex
	//TODO: Database connection
}

type ModuleJSON struct {
	Author      string            `json:"author"`
	DisplayName map[string]string `json:"displayName"`
}

type ModuleFactory func(context ModuleContext) (Module, error)
