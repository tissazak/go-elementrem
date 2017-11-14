// Copyright 2016-2017 The go-elementrem Authors
// This file is part of the go-elementrem library.
//
// The go-elementrem library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-elementrem library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-elementrem library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"

	"github.com/tissazak/go-elementrem/common"
)

var (
	TestNetGenesisHash = common.HexToHash("") // Testnet genesis hash to enforce below configs on
	MainNetGenesisHash = common.HexToHash("") // Mainnet genesis hash to enforce below configs on

	TestNetHomesteadBlock = big.NewInt(0)       // Testnet homestead block
	MainNetHomesteadBlock = big.NewInt(0) // Mainnet homestead block

	TestNetHomesteadGasRepriceBlock = big.NewInt(0)       // Testnet gas reprice block
	MainNetHomesteadGasRepriceBlock = big.NewInt(1900000) // Mainnet gas reprice block

	TestNetHomesteadGasRepriceHash = common.HexToHash("") // Testnet gas reprice block hash (used by fast sync)
	MainNetHomesteadGasRepriceHash = common.HexToHash("0") // Mainnet gas reprice block hash (used by fast sync)

	TestNetSpuriousDragon = big.NewInt(0)
	MainNetSpuriousDragon = big.NewInt(1900000)

	MainNetTronRecursive = big.NewInt(1900000) 

	TestNetChainID = big.NewInt(70709) // Test net default chain ID
	MainNetChainID = big.NewInt(1989) // main net default chain ID
)
