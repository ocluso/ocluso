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

package core

import (
	"context"
	"database/sql"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ocluso/ocluso/backend/modules/accounts"
	"github.com/ocluso/ocluso/backend/modules/members"
)

// Server serves the ocluso backend
type Server struct {
	db           *sql.DB
	httpListener *net.Listener
	httpServer   *http.Server
}

// NewServer creates a new ocluso backend server that conforms to the given configuration
//
// The database connection will be established during this call
func NewServer(config Config) (*Server, error) {
	listener, err := net.Listen("tcp", config.HTTPListenAddress)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", config.PostgresDSN)
	if err != nil {
		return nil, err
	}

	server := http.Server{}
	router := mux.NewRouter()
	addHandlersForModule(router, "accounts", accounts.BuildHandler(db))
	addHandlersForModule(router, "members", members.BuildHandler(db))
	server.Handler = router

	return &Server{db: db, httpListener: &listener, httpServer: &server}, nil
}

// Serve runs the ocluso backend server in the current goroutine
func (s *Server) Serve() error {
	return s.httpServer.Serve(*s.httpListener)
}

// Shutdown gracefully shuts down the ocluso backend server
func (s *Server) Shutdown() error {
	if err := s.db.Close(); err != nil {
		return err
	}

	return s.httpServer.Shutdown(context.Background())
}

func buildPathPrefix(moduleName string) string {
	return "/api/v0/" + moduleName
}

func addHandlersForModule(router *mux.Router, moduleName string, handler http.Handler) {
	prefix := buildPathPrefix(moduleName)

	router.Path(prefix).Handler(http.StripPrefix(prefix, handler))
	router.PathPrefix(prefix + "/").Handler(http.StripPrefix(prefix, handler))
}
