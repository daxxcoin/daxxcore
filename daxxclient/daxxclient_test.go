// Copyright 2016 The daxxcoreAuthors
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

package ethclient

import "github.com/daxxcoin/daxxcore"

// Verify that Client implements the daxxcoin interfaces.
var (
	_ = daxxcoin.ChainReader(&Client{})
	_ = daxxcoin.TransactionReader(&Client{})
	_ = daxxcoin.ChainStateReader(&Client{})
	_ = daxxcoin.ChainSyncReader(&Client{})
	_ = daxxcoin.ContractCaller(&Client{})
	_ = daxxcoin.GasEstimator(&Client{})
	_ = daxxcoin.GasPricer(&Client{})
	_ = daxxcoin.LogFilterer(&Client{})
	_ = daxxcoin.PendingStateReader(&Client{})
	// _ = daxxcoin.PendingStateEventer(&Client{})
	_ = daxxcoin.PendingContractCaller(&Client{})
)
