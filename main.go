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

	"github.com/ocluso/ocluso/internal/gateway"
)

func main() {
	moduleGateway := gateway.NewGateway()

	//TODO: Auto-detect module names or load from configuration file
	moduleGateway.AddModule("calendar")
	moduleGateway.AddModule("fees")
	moduleGateway.AddModule("mailinglists")
	moduleGateway.AddModule("members")

	http.Handle("/", http.FileServer(http.Dir("frontend"))) //TODO: Compile frontend into backend
	http.Handle("/api/", &moduleGateway)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
