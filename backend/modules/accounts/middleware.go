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
	"github.com/gorilla/mux"
	"github.com/lhinderberger/KISStokens"
)

// BuildAuthenticationMiddleware builds a middleware for mux.Router that checks for authentication tokens,
// verifies them and, rejects the request if invalid and otherwise injects the decoded values of the
// authentication token into the request's context.
func BuildAuthenticationMiddleware(tokenAuthority *KISStokens.TokenAuthority) mux.MiddlewareFunc {
	panic("Not implemented")
}
