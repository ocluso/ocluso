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
	"github.com/lhinderberger/KISStokens"
	"net/http"

	"github.com/gorilla/mux"
)

func BuildHandler(db *sql.DB, tokenAuthority *KISStokens.TokenAuthority) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/login", buildLoginHandler(db, tokenAuthority))
	router.HandleFunc("/logout", buildLogoutHandler())
	router.HandleFunc("/me", buildMeHandler(db))

	return router
}