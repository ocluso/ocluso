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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPasswordYieldsCorrectHash(t *testing.T) {
	password := "FooBarBaz"

	hash, err := HashPassword(password)
	assert.NoError(t, err)

	verified, err := VerifyPassword(password, hash)
	assert.NoError(t, err)
	assert.True(t, verified)
}

func TestHashPasswordDoesntYieldSameHashTwice(t *testing.T) {
	password := "FooBarBaz"

	hash1, err := HashPassword(password)
	assert.NoError(t, err)

	hash2, err := HashPassword(password)
	assert.NoError(t, err)

	assert.NotEqual(t, hash1, hash2)
}

func TestVerifyPasswordDetectsFalsePassword(t *testing.T) {
	password := "FooBarBaz"

	hash, err := HashPassword(password)
	assert.NoError(t, err)

	verified, err := VerifyPassword("NotThePassword", hash)
	assert.NoError(t, err)
	assert.False(t, verified)
}

func TestVerifyPasswordDetectsInvalidHashLength(t *testing.T) {
	_, err := VerifyPassword("NotThePassword", "NotAValidHash")
	assert.Error(t, err)

	_, err = VerifyPassword("NotThePassword", strings.Repeat("X", 129))
	assert.Error(t, err)
}

func TestVerifyPasswordDetectsInvalidHashValue(t *testing.T) {
	_, err := VerifyPassword("NotThePassword", strings.Repeat("X", 128))
	assert.Error(t, err)
}
