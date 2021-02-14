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
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lhinderberger/KISStokens"
	"github.com/ocluso/ocluso/backend/testutils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func buildMiddlewareTestRequest(target string, authToken string, csrfToken string) *http.Request {
	body := fmt.Sprintf("Cookie: %s=%s\n%s: %s", AuthTokenCookieName, authToken, CSRFHeaderName, csrfToken)
	return httptest.NewRequest("POST", target, strings.NewReader(body))
}

func buildMiddlewareTestRequestNoCSRF(target string, authToken string) *http.Request {
	body := fmt.Sprintf("Cookie: %s=%s", AuthTokenCookieName, authToken)
	return httptest.NewRequest("POST", target, strings.NewReader(body))
}

func buildMiddlewareTestRouter(tokenAuthority *KISStokens.TokenAuthority) *mux.Router {
	middleware := BuildAuthenticationMiddleware(tokenAuthority)

	router := mux.NewRouter()
	router.Use(middleware)
	router.Handle("/foo", http.HandlerFunc(fooHandler))
	router.Handle("/requireUserAuthenticated", http.HandlerFunc(requireUserAuthenticatedHandler))

	return router
}

func fooHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("foo"))
}

func requireUserAuthenticatedHandler(w http.ResponseWriter, r *http.Request) {
	userUUID, ok := RequireUserAuthenticated(w, r)
	if ok {
		w.Write([]byte(userUUID))
	}
}

//TODO: Add integration tests for all paths
func TestValidAuthTokenInjectsContext(t *testing.T) {
	tokenAuthority := buildTestTokenAuthority()
	router := buildMiddlewareTestRouter(tokenAuthority)

	csrfToken := "00112233445566778899aabbccddeeff00112233445566778899aabbccddeeff"
	expectedUserUUID := "6d4ea271-97e1-42f6-9187-ae28ce5a4cd0"
	authToken, err := tokenAuthority.Sign(KISStokens.Claims{
		"csrf": csrfToken,
		"user": expectedUserUUID,
	})
	if !assert.NoError(t, err) {
		t.Fatal("Could not continue test")
	}

	request := buildMiddlewareTestRequest("/requireUserAuthenticated", authToken, csrfToken)
	response := testutils.RecordResponse(router, request)

	if assert.Equal(t, http.StatusOK, response.Status) {
		actualUserUUID, err := ioutil.ReadAll(response.Body)
		if assert.NoError(t, err) {
			assert.Equal(t, expectedUserUUID, string(actualUserUUID))
		}
	}
}

func TestInjectAuthenticationInjectsContext(t *testing.T) {
	expectedUserUUID := "735fa28e-3cb0-4a73-ab8b-a0dda7402151"
	handler := InjectAuthentication(expectedUserUUID, http.HandlerFunc(requireUserAuthenticatedHandler))

	request := httptest.NewRequest("POST", "/requireUserAuthenticated", nil)
	response := testutils.RecordResponse(handler, request)

	if assert.Equal(t, http.StatusOK, response.Status) {
		actualUserUUID, err := ioutil.ReadAll(response.Body)
		if assert.NoError(t, err) {
			assert.Equal(t, expectedUserUUID, string(actualUserUUID))
		}
	}
}

func TestInvalidAuthTokenYields401(t *testing.T) {
	tokenAuthority := buildTestTokenAuthority()
	router := buildMiddlewareTestRouter(tokenAuthority)

	csrfToken := "00112233445566778899aabbccddeeff00112233445566778899aabbccddeeff"
	expectedUserUUID := "6d4ea271-97e1-42f6-9187-ae28ce5a4cd0"
	authToken, err := tokenAuthority.SignWithLimits(KISStokens.Claims{
		"csrf": csrfToken,
		"user": expectedUserUUID,
	}, KISStokens.Limits{
		ExpirationTime: time.Now().Add(-time.Minute),
		NotBefore:      KISStokens.NoTime(),
	})
	if !assert.NoError(t, err) {
		t.Fatal("Could not continue test")
	}

	request := buildMiddlewareTestRequest("/requireUserAuthenticated", authToken, csrfToken)
	response := testutils.RecordResponse(router, request)

	assert.Equal(t, http.StatusUnauthorized, response.Status)
}

func TestMissingCSRFYields401(t *testing.T) {
	tokenAuthority := buildTestTokenAuthority()
	router := buildMiddlewareTestRouter(tokenAuthority)

	csrfToken := "00112233445566778899aabbccddeeff00112233445566778899aabbccddeeff"
	expectedUserUUID := "6d4ea271-97e1-42f6-9187-ae28ce5a4cd0"
	authToken, err := tokenAuthority.Sign(KISStokens.Claims{
		"csrf": csrfToken,
		"user": expectedUserUUID,
	})
	if !assert.NoError(t, err) {
		t.Fatal("Could not continue test")
	}

	request := buildMiddlewareTestRequestNoCSRF("/requireUserAuthenticated", authToken)
	response := testutils.RecordResponse(router, request)

	assert.Equal(t, http.StatusUnauthorized, response.Status)
}

func TestInvalidCSRFYields401(t *testing.T) {
	tokenAuthority := buildTestTokenAuthority()
	router := buildMiddlewareTestRouter(tokenAuthority)

	csrfToken := "00112233445566778899aabbccddeeff00112233445566778899aabbccddeeff"
	expectedUserUUID := "6d4ea271-97e1-42f6-9187-ae28ce5a4cd0"
	authToken, err := tokenAuthority.Sign(KISStokens.Claims{
		"csrf": csrfToken,
		"user": expectedUserUUID,
	})
	if !assert.NoError(t, err) {
		t.Fatal("Could not continue test")
	}

	request := buildMiddlewareTestRequest("/requireUserAuthenticated", authToken, "foo")
	response := testutils.RecordResponse(router, request)

	assert.Equal(t, http.StatusUnauthorized, response.Status)
}

func TestNoAuthTokenIsANop(t *testing.T) {
	tokenAuthority := buildTestTokenAuthority()
	router := buildMiddlewareTestRouter(tokenAuthority)

	request := httptest.NewRequest("POST", "/foo", nil)
	response := testutils.RecordResponse(router, request)

	if assert.Equal(t, http.StatusOK, response.Status) {
		actualFoo, err := ioutil.ReadAll(response.Body)
		if assert.NoError(t, err) {
			assert.Equal(t, "foo", string(actualFoo))
		}
	}
}
