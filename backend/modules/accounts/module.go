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

package accounts

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func BuildHandler(db *sql.DB) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from accounts/foo"))
	})

	router.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from accounts/bar"))
	})

	return router
}
