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

package testutils

import (
	"net/http"
	"net/http/httptest"
)

func ExtractCookieFromResponse(response *http.Response, cookieName string) *http.Cookie {
	for _, cookie := range response.Cookies() {
		if cookie.Name == cookieName {
			return cookie
		}
	}

	return nil
}

// RecordResponse records an http response to a given request for a given handler
func RecordResponse(handler http.Handler, request *http.Request) *http.Response {
	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	return responseRecorder.Result()
}
