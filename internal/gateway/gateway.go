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

package gateway

import "net/http"

type Gateway struct {
	//TODO
}

func NewGateway() Gateway {
	return Gateway{}
}

func (g *Gateway) AddModule(name string) {
	//panic("Not implemented")
}

func (g *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gateway: Not implemented!"))
}
