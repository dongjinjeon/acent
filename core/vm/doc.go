// Copyright 2015 The go-acent Authors
// This file is part of the go-acent library.
//
// The go-acent library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-acent library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-acent library. If not, see <http://www.gnu.org/licenses/>.

/*
Package vm implements the Acent Virtual Machine.

The vm package implements one EVM, a byte code VM. The BC (Byte Code) VM loops
over a set of bytes and executes them according to the set of rules defined
in the Acent yellow paper.
*/
package vm
