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

package main

import (
	"log"
	"net/http"
	"strings"

	"ocluso/internal/gateway"
	"ocluso/pkg/moduleindex"
	"ocluso/pkg/moduleinterface"
)

func main() {
	printLoadedModules()

	moduleGateway := gateway.NewGateway()

	for _, loadedModule := range moduleindex.LoadedModules {
		context := moduleinterface.ModuleContext{}
		module, err := loadedModule.Factory(context)
		if err != nil {
			panic(err)
		}

		moduleGateway.AddModule(module)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("frontend"))) //TODO: Compile frontend into backend
	mux.Handle("/api/", http.StripPrefix("/api", &moduleGateway))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func printLoadedModules() {
	moduleNames := make([]string, len(moduleindex.LoadedModules))
	i := 0
	for name := range moduleindex.LoadedModules {
		moduleNames[i] = name
		i++
	}

	log.Println("Loaded Modules:", strings.Join(moduleNames, ", "))
}
