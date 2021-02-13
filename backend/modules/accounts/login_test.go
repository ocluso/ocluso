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
	"encoding/base64"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lhinderberger/KISStokens"
	"github.com/ocluso/ocluso/backend/testutils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const loginTestUsername = "example@ocluso.de"

func assertAuthTokenIsValid(t *testing.T, tokenAuthority *KISStokens.TokenAuthority, authTokenCookie *http.Cookie, csrfHeader string) {
	if assert.NotNil(t, authTokenCookie) && assert.NotEmpty(t, csrfHeader) {
		assert.True(t, authTokenCookie.HttpOnly)
		assert.True(t, authTokenCookie.Secure)
		assert.Equal(t, http.SameSiteStrictMode, authTokenCookie.SameSite)

		decodedToken, err := tokenAuthority.DecodeAndVerify(authTokenCookie.Value)
		assert.NoError(t, err)

		csrfTokenField, ok := decodedToken.Claims[CSRFClaimName]
		if assert.True(t, ok) {
			assert.Equal(t, csrfHeader, csrfTokenField)
		}

		assert.Equal(t, decodedToken.Header.Limits.ExpirationTime, authTokenCookie.Expires)
	}
}

func buildTestTokenAuthority() *KISStokens.TokenAuthority {
	key := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}

	authority, err := KISStokens.NewTokenAuthority(key, 30*24*time.Hour)
	if err != nil {
		panic(fmt.Sprint("Could not build test token authority:", err))
	}

	return authority
}

func buildLoginTestMockDatabase() (*sql.DB, *sqlmock.Sqlmock) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		panic("Cannot create mock database")
	}

	rows := sqlmock.
		NewRows([]string{"Accounts.passwordHash"}).
		AddRow("d1dc6e1dec083a369fe747aba354daa20430e941c663590f6b628daf68685ed7e4692583b68b66a16a052ac4acd30710b976a84b4b14dd404ac7b17f84e56a2d")

	dbMock.ExpectQuery(passwordHashQuery).WithArgs(loginTestUsername).WillReturnRows(rows)

	return db, &dbMock
}

func buildGenericLoginTestRequest(body string) *http.Request {
	return httptest.NewRequest("POST", "/accounts/login", strings.NewReader(body))
}

func buildLoginTestRequest(email string, password string) *http.Request {
	authHeaderPayload := fmt.Sprintf("%s:%s", email, password)
	body := "Authorization: Basic " + base64.StdEncoding.EncodeToString([]byte(authHeaderPayload))
	return buildGenericLoginTestRequest(body)
}

func buildLoginTestRenewRequest(authTokenCookie *http.Cookie, csrfToken string) *http.Request {
	body := fmt.Sprintf("Cookie: %s=%s\n%s: %s", authTokenCookie.Name, authTokenCookie.Value, CSRFHeaderName, csrfToken)
	return buildGenericLoginTestRequest(body)
}

func buildLogoutTestRequest() *http.Request {
	return httptest.NewRequest("POST", "/accounts/logout", nil)
}

//TODO: Add integration test for happy path and invalid password
//TODO: Add integration test for deleted users
//TODO: Add integration test for possible timing attacks
func TestLoginWithValidCredentialsYieldsValidAuthToken(t *testing.T) {
	db, dbMock := buildLoginTestMockDatabase()
	tokenAuthority := buildTestTokenAuthority()
	handler := buildLoginHandler(db, tokenAuthority)

	request := buildLoginTestRequest(loginTestUsername, "FooBarBaz")
	response := testutils.RecordResponse(handler, request)
	if !assert.Equal(t, http.StatusOK, response.StatusCode) {
		t.Fatal("Cannot complete test")
	}

	assert.NoError(t, (*dbMock).ExpectationsWereMet())

	authTokenCookie := testutils.ExtractCookieFromResponse(response, AuthTokenCookieName)
	csrfHeader := response.Header.Get(CSRFHeaderName)

	assertAuthTokenIsValid(t, tokenAuthority, authTokenCookie, csrfHeader)
}

func TestLoginWithValidTokenRenewsToken(t *testing.T) {
	db, _ := buildLoginTestMockDatabase()
	tokenAuthority := buildTestTokenAuthority()
	handler := buildLoginHandler(db, tokenAuthority)

	request := buildLoginTestRequest(loginTestUsername, "FooBarBaz")
	response := testutils.RecordResponse(handler, request)
	if !assert.Equal(t, http.StatusOK, response.StatusCode) {
		t.Fatal("Cannot complete test")
	}

	authTokenCookie := testutils.ExtractCookieFromResponse(response, AuthTokenCookieName)
	csrfHeader := response.Header.Get(CSRFHeaderName)

	time.Sleep(time.Second)

	request = buildLoginTestRenewRequest(authTokenCookie, csrfHeader)
	response = testutils.RecordResponse(handler, request)

	if !assert.Equal(t, http.StatusOK, response.StatusCode) {
		t.Fatal("Cannot complete test")
	}

	authTokenCookie2 := testutils.ExtractCookieFromResponse(response, AuthTokenCookieName)
	csrfHeader2 := response.Header.Get(CSRFHeaderName)

	assertAuthTokenIsValid(t, tokenAuthority, authTokenCookie2, csrfHeader2)

	assert.NotEqual(t, authTokenCookie, authTokenCookie2)
	assert.NotEqual(t, csrfHeader, csrfHeader2)
	assert.True(t, authTokenCookie2.Expires.After(authTokenCookie.Expires))
}

func TestLoginWithNonExistingUsernameYields401(t *testing.T) {
	db, _ := buildLoginTestMockDatabase()
	tokenAuthority := buildTestTokenAuthority()
	handler := buildLoginHandler(db, tokenAuthority)

	request := buildLoginTestRequest("doesntexist@ocluso.de", "FooBarBaz")
	response := testutils.RecordResponse(handler, request)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestLoginWithIncorrectPasswordYields401(t *testing.T) {
	db, _ := buildLoginTestMockDatabase()
	tokenAuthority := buildTestTokenAuthority()
	handler := buildLoginHandler(db, tokenAuthority)

	request := buildLoginTestRequest(loginTestUsername, "SomeOtherPassword")
	response := testutils.RecordResponse(handler, request)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestLogoutDropsCookie(t *testing.T) {
	handler := buildLogoutHandler()

	request := buildLogoutTestRequest()
	response := testutils.RecordResponse(handler, request)

	if !assert.Equal(t, http.StatusOK, response.StatusCode) {
		t.Fatal("Cannot complete test")
	}

	authTokenCookie := testutils.ExtractCookieFromResponse(response, AuthTokenCookieName)

	if assert.NotNil(t, authTokenCookie) {
		assert.Empty(t, authTokenCookie.Value)
		assert.True(t, authTokenCookie.Expires.Before(time.Now()))
	}
}
