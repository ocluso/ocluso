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

CREATE TABLE Genders (
    id SERIAL PRIMARY KEY,
    genderName VARCHAR(32) NOT NULL,
    genderNamePlural VARCHAR(32)
);

CREATE TABLE MemberTypes (
    id SERIAL PRIMARY KEY,
    typeName VARCHAR(64) NOT NULL,
    typeNamePlural VARCHAR(64)
);

CREATE TABLE Members (
    uuid UUID PRIMARY KEY,
    clubMemberId VARCHAR(64) UNIQUE,
    memberType SERIAL NOT NULL REFERENCES MemberTypes,
    firstName VARCHAR(64) NOT NULL,
    lastName VARCHAR(64) NOT NULL,
    birthName VARCHAR(64),
    birthday DATE NOT NULL,
    street VARCHAR(64) NOT NULL,
    houseNumber VARCHAR(16) NOT NULL,
    addition VARCHAR(64) NOT NULL,
    zip VARCHAR(16) NOT NULL,
    town VARCHAR(64) NOT NULL,
    countryCode CHAR(2) NOT NULL,
    email VARCHAR(256),
    phone VARCHAR(32),
    dateEntered DATE NOT NULL,
    dateLeft DATE,
    gender SERIAL NOT NULL REFERENCES Genders
);
