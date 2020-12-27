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

import (
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fooModule struct{}

type barModule struct{}

const fooMessage = "Hello Foo"

func (m *fooModule) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte(fooMessage))
}

func (m *barModule) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte(request.URL.Path))
}

func buildTestListener() net.Listener {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	return listener
}

func buildTestServer() *http.Server {
	modules := map[string]Module{
		"foo": &fooModule{},
		"bar": &barModule{},
	}

	server := http.Server{}
	server.Handler = NewModuleMux(&modules)

	return &server
}

func testRouteLeadsTo404(t *testing.T, route string) {
	listener := buildTestListener()
	server := buildTestServer()

	go server.Serve(listener)
	defer server.Close()

	response, err := http.Get("http://" + listener.Addr().String() + "/" + route)
	assert.NoError(t, err)
	assert.Equal(t, 404, response.StatusCode)
}

func testRouteYieldsBody(t *testing.T, route string, expectedBody string) {
	expectedBytes := []byte(expectedBody)
	listener := buildTestListener()
	server := buildTestServer()

	go server.Serve(listener)
	defer server.Close()

	response, err := http.Get("http://" + listener.Addr().String() + "/" + route)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)
	assert.EqualValues(t, len(expectedBytes), response.ContentLength)

	responseBody := make([]byte, response.ContentLength)
	_, err = io.ReadFull(response.Body, responseBody)
	assert.NoError(t, err)
	assert.Equal(t, expectedBytes, responseBody)
}

func TestModuleIsAvailableAtSubpath(t *testing.T) {
	testRouteYieldsBody(t, "foo", fooMessage)
}

func TestNonExistingModuleYields404(t *testing.T) {
	testRouteLeadsTo404(t, "baz")
}

func TestRouteYields404(t *testing.T) {
	testRouteLeadsTo404(t, "")
}

func TestPrefixIsStripped(t *testing.T) {
	testRouteYieldsBody(t, "bar/baz", "baz")
}
