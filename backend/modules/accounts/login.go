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
)

const AuthTokenCookieName = "authToken"
const CSRFClaimName = "csrf"
const CSRFHeaderName = "X-CSRF-Token"

const passwordHashQuery = "SELECT Accounts.passwordHash FROM Accounts JOIN Members on Members.memberUUID = Accounts.memberUUID WHERE Members.email = ?"

func buildLoginHandler(db *sql.DB, tokenAuthority *KISStokens.TokenAuthority) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from accounts/login"))

		//TODO: Look up submitted user
		//TODO: Check submitted password
		//TODO: Issue authentication token, if password is correct
		//TODO: Send that token using secure cookie
		//TODO: Send CSRF token as header
	})
}

func buildLogoutHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from accounts/logout"))

		//TODO: Invalidate authentication cookie, if any
	})
}
