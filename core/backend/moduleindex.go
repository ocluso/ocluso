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

type ModuleIndex struct {
	//TODO
}

type ModuleIndexBuilder struct {
	//TODO
}

type ModuleIndexEntry struct {
	ModuleJson    ModuleJson
	ModuleFactory ModuleFactory
}

func NewModuleIndexBuilder() *ModuleIndexBuilder {
	return &ModuleIndexBuilder{}
}

func (b *ModuleIndexBuilder) AddModule(name string, entry *ModuleIndexEntry) error {
	panic("Not implemented")
}

func (b *ModuleIndexBuilder) Build() *ModuleIndex {
	panic("Not implemented")
}

func (m *ModuleIndex) EntryFor(moduleName string) (*ModuleIndexEntry, error) {
	panic("Not implemented")
}

func (m *ModuleIndex) LoadedModules() []string {
	panic("Not implemented")
}
