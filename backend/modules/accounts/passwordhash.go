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
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/argon2"
)

// HashPassword hashes a password for use with ocluso
// using Argon2, producing a 128 character long hash
func HashPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	hash := hashWithSalt(password, salt)

	return (hex.EncodeToString(salt) + hex.EncodeToString(hash)), nil
}

// VerifyPassword checks a password against a given hash
// that was created using HashPassword
//
// It returns true, if the password equals the one that was hashed,
// otherwise false.
func VerifyPassword(password string, hash string) (bool, error) {
	if len(hash) != 128 {
		return false, errors.New("Cannot verify hash: Invalid hash length")
	}

	decodedSaltAndHash, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	decodedSalt := decodedSaltAndHash[:32]
	decodedHash := decodedSaltAndHash[32:]

	computedHash := hashWithSalt(password, decodedSalt)

	return bytes.Equal(decodedHash, computedHash), nil
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	return salt, err
}

func hashWithSalt(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, 1, 0xFFFF, 4, 32)
}
