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
	"errors"
	"fmt"
)

// ModuleIndex is an index of (non-instantiated) ocluso Modules
//
// It contains general information about and factory functions for ocluso
// Modules that were compiled into the current ocluso binary
type ModuleIndex struct {
	entries map[string]ModuleIndexEntry
}

// ModuleIndexEntry contains information about an available ocluso Module
type ModuleIndexEntry struct {
	ModuleJSON    ModuleJSON
	ModuleFactory ModuleFactory
}

// NewModuleIndex creates a new module index from the given entries map
func NewModuleIndex(entries *map[string]ModuleIndexEntry) *ModuleIndex {
	return &ModuleIndex{entries: *entries}
}

// EntryFor returns the entry for a given module name
//
// If there is no module with the given name in the index, an error is returned
func (m *ModuleIndex) EntryFor(moduleName string) (*ModuleIndexEntry, error) {
	entry, contained := m.entries[moduleName]

	if contained {
		return &entry, nil
	}

	return nil, errors.New(fmt.Sprint("No such module:", moduleName))
}

// LoadedModules returns a list of the names of all modules in the index
func (m *ModuleIndex) LoadedModules() []string {
	moduleNames := make([]string, len(m.entries))

	i := 0
	for moduleName := range m.entries {
		moduleNames[i] = moduleName
		i++
	}

	return moduleNames
}
