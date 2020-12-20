# Copyright (C) 2020 The ocluso Authors
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.


build: build_dirs moduleindex frontend
	go build -o build/ocluso main.go 

build_dirs:
	mkdir -p build/frontend

frontend: build_dirs moduleindex
	echo "Hello World from frontend!" > build/frontend/index.html

moduleindex:
	go run tools/gen-moduleindex/gen-moduleindex.go

test: test-backend test-frontend

test-backend:
	go test ./...

test-frontend: moduleindex
	cd frontend && npm test

clean:
	rm -r build

run: build
	cd build && ./ocluso
