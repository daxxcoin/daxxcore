// Copyright 2014 The daxxcoreAuthors
// This file is part of the daxxcore library.
//
// The daxxcore library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The daxxcore library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the daxxcore library. If not, see <http://www.gnu.org/licenses/>.

package vm

import "errors"

var (
	ErrOutOfGas            = errors.New("out of gas")
	ErrCodeStoreOutOfGas   = errors.New("contract creation code storage out of gas")
	ErrDepth               = errors.New("max call depth exceeded")
	ErrTraceLimitReached   = errors.New("the number of logs reached the specified limit")
	ErrInsufficientBalance = errors.New("insufficient balance for transfer")
)