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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadedModulesYieldsCorrectList(t *testing.T) {
	expectedModuleNames := []string{"foo", "bar", "baz"}

	entries := make(map[string]ModuleIndexEntry, len(expectedModuleNames))
	for _, moduleName := range expectedModuleNames {
		entries[moduleName] = ModuleIndexEntry{}
	}

	index := NewModuleIndex(&entries)

	actualModuleNames := index.LoadedModules()

	assert.ElementsMatch(t, expectedModuleNames, actualModuleNames)
}

func TestEntryForYieldsCorrectEntry(t *testing.T) {
	expectedAuthor := "John Doe"
	entries := map[string]ModuleIndexEntry{
		"foo": ModuleIndexEntry{ModuleJSON: ModuleJSON{Author: expectedAuthor}},
		"bar": ModuleIndexEntry{},
	}

	index := NewModuleIndex(&entries)

	entry, err := index.EntryFor("foo")

	assert.NoError(t, err)
	assert.Equal(t, entry.ModuleJSON.Author, expectedAuthor)
}

func TestEntryForYieldsErrorForNonExistingModule(t *testing.T) {
	index := NewModuleIndex(&map[string]ModuleIndexEntry{"foo": ModuleIndexEntry{}})

	_, err := index.EntryFor("bar")
	assert.Error(t, err)
}
