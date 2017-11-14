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

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Elementrem network.
var MainnetBootnodes = []string{
	// ELE/DEV Go Bootnodes
	"enode://bb4ede6faadc19749e9119bcb8c487e10c2651ffa0a4aaf62e89431d133cc12d9bb8ba3858a10fd9f2e0f961b8db447ff8d2eaa2f962014729ebcff86f8f8d7f@35.177.83.134:30303",
	"enode://0a946018428af2188b3fbb11490c19f10bf1f8b868862a2fff2e1c1287ccb5bb3296d12095247e9e19a6a3d7eee6f98647e928c129322c0b816c9ae79ef86b84@35.177.85.183:30303"
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestnetBootnodes = []string{
	// ELE/DEV Go Bootnodes
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
