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

import "testing"

//TODO: Add integration test for happy path
func TestLoginWithValidCredentialsYieldsAuthToken(t *testing.T) {
	t.Fatal("Not implemented")
}

func TestLoginWithValidTokenRenewsToken(t *testing.T) {
	t.Fatal("Not implemented")
}

func TestLoginWithNonExistingUsernameYields401(t *testing.T) {
	t.Fatal("Not implemented")
}

func TestLoginWithDeletedUserYields401(t *testing.T) {
	t.Fatal("Not implemented")
}

func TestLoginWithIncorrectPasswordYields401(t *testing.T) {
	t.Fatal("Not implemented")
}

func TestLoginTakesTheConfiguredTime(t *testing.T) {
	t.Fatal("Not implemented")
}

func TestLogoutDropsCookie(t *testing.T) {
	t.Fatal("Not implemented")
}

func TestLogoutWithoutValidTokenAlsoDropsCookie(t *testing.T) {
	t.Fatal("Not implemented")
}
