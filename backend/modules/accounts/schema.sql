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

CREATE TYPE Permission AS (
    moduleName VARCHAR(64)
    permissionName VARCHAR(32)
);


CREATE TABLE Accounts (
    memberUUID UUID PRIMARY KEY REFERENCES Members ON DELETE CASCADE,
    passwordHash CHAR(128) NOT NULL,
    isAdmin BOOLEAN NOT NULL DEFAULT false,
    profilePicture BYTEA
);

CREATE TABLE Invites (
    memberUUID UUID PRIMARY KEY REFERENCES Members ON DELETE CASCADE,
    randomTokenHash CHAR(128) NOT NULL,
    issued TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE PermissionGroups (
    id SERIAL PRIMARY KEY,
    groupName VARCHAR(64) NOT NULL
);


CREATE TABLE AccountPermissions (
    memberUUID UUID NOT NULL REFERENCES Accounts ON DELETE CASCADE,
    permission Permission NOT NULL,
    PRIMARY KEY(memberUUID, permission)
);

CREATE TABLE AccountPermissionGroups (
    memberUUID UUID NOT NULL REFERENCES Accounts ON DELETE CASCADE,
    permissionGroupId SERIAL NOT NULL REFERENCES PermissionGroups ON DELETE CASCADE,
    PRIMARY KEY(memberUUID, permissionGroupId)
);

CREATE TABLE PermissionGroupPermissions (
    permissionGroupId SERIAL NOT NULL REFERENCES PermissionGroups ON DELETE CASCADE,
    permission Permission NOT NULL,
    PRIMARY KEY(permissionGroupId, permission)
);
