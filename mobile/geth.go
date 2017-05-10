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

// Contains all the wrappers from the node package to support client side node
// management on mobile platforms.

package geth

import (
	"fmt"
	"math/big"
	"path/filepath"

	"github.com/daxxcoin/daxxcore/common"
	"github.com/daxxcoin/daxxcore/daxx"
	"github.com/daxxcoin/daxxcore/daxxclient"
	"github.com/daxxcoin/daxxcore/daxxstats"
	"github.com/daxxcoin/daxxcore/les"
	"github.com/daxxcoin/daxxcore/node"
	"github.com/daxxcoin/daxxcore/p2p/nat"
	"github.com/daxxcoin/daxxcore/params"
	"github.com/daxxcoin/daxxcore/whisper/whisperv2"
)

// NodeConfig represents the collection of configuration values to fine tune the Geth
// node embedded into a mobile process. The available values are a subset of the
// entire API provided by daxxcoreto reduce the maintenance surface and dev
// complexity.
type NodeConfig struct {
	// Bootstrap nodes used to establish connectivity with the rest of the network.
	BootstrapNodes *Enodes

	// MaxPeers is the maximum number of peers that can be connected. If this is
	// set to zero, then only the configured static and trusted peers can connect.
	MaxPeers int

	// DaxxcoinEnabled specifies whether the node should run the Daxxcoin protocol.
	DaxxcoinEnabled bool

	// DaxxcoinNetworkID is the network identifier used by the Daxxcoin protocol to
	// decide if remote peers should be accepted or not.
	DaxxcoinNetworkID int

	// DaxxcoinChainConfig is the default parameters of the blockchain to use. If no
	// configuration is specified, it defaults to the main network.
	DaxxcoinChainConfig *ChainConfig

	// DaxxcoinGenesis is the genesis JSON to use to seed the blockchain with. An
	// empty genesis state is equivalent to using the mainnet's state.
	DaxxcoinGenesis string

	// DaxxcoinDatabaseCache is the system memory in MB to allocate for database caching.
	// A minimum of 16MB is always reserved.
	DaxxcoinDatabaseCache int

	// DaxxcoinNetStats is a netstats connection string to use to report various
	// chain, transaction and node stats to a monitoring server.
	//
	// It has the form "nodename:secret@host:port"
	DaxxcoinNetStats string

	// WhisperEnabled specifies whether the node should run the Whisper protocol.
	WhisperEnabled bool
}

// defaultNodeConfig contains the default node configuration values to use if all
// or some fields are missing from the user's specified list.
var defaultNodeConfig = &NodeConfig{
	BootstrapNodes:        FoundationBootnodes(),
	MaxPeers:              25,
	DaxxcoinEnabled:       true,
	DaxxcoinNetworkID:     1,
	DaxxcoinChainConfig:   MainnetChainConfig(),
	DaxxcoinDatabaseCache: 16,
}

// NewNodeConfig creates a new node option set, initialized to the default values.
func NewNodeConfig() *NodeConfig {
	config := *defaultNodeConfig
	return &config
}

// Node represents a Geth Daxxcoin node instance.
type Node struct {
	node *node.Node
}

// NewNode creates and configures a new Geth node.
func NewNode(datadir string, config *NodeConfig) (stack *Node, _ error) {
	// If no or partial configurations were specified, use defaults
	if config == nil {
		config = NewNodeConfig()
	}
	if config.MaxPeers == 0 {
		config.MaxPeers = defaultNodeConfig.MaxPeers
	}
	if config.BootstrapNodes == nil || config.BootstrapNodes.Size() == 0 {
		config.BootstrapNodes = defaultNodeConfig.BootstrapNodes
	}
	// Create the empty networking stack
	nodeConf := &node.Config{
		Name:             clientIdentifier,
		Version:          params.Version,
		DataDir:          datadir,
		KeyStoreDir:      filepath.Join(datadir, "keystore"), // Mobile should never use internal keystores!
		NoDiscovery:      true,
		DiscoveryV5:      true,
		DiscoveryV5Addr:  ":0",
		BootstrapNodesV5: config.BootstrapNodes.nodes,
		ListenAddr:       ":0",
		NAT:              nat.Any(),
		MaxPeers:         config.MaxPeers,
	}
	rawStack, err := node.New(nodeConf)
	if err != nil {
		return nil, err
	}
	// Register the Daxxcoin protocol if requested
	if config.DaxxcoinEnabled {
		ethConf := &eth.Config{
			ChainConfig: &params.ChainConfig{
				ChainId:        big.NewInt(config.DaxxcoinChainConfig.ChainID),
				HomesteadBlock: big.NewInt(config.DaxxcoinChainConfig.HomesteadBlock),
				DAOForkBlock:   big.NewInt(config.DaxxcoinChainConfig.DAOForkBlock),
				DAOForkSupport: config.DaxxcoinChainConfig.DAOForkSupport,
				EIP150Block:    big.NewInt(config.DaxxcoinChainConfig.EIP150Block),
				EIP150Hash:     config.DaxxcoinChainConfig.EIP150Hash.hash,
				EIP155Block:    big.NewInt(config.DaxxcoinChainConfig.EIP155Block),
				EIP158Block:    big.NewInt(config.DaxxcoinChainConfig.EIP158Block),
			},
			Genesis:                 config.DaxxcoinGenesis,
			LightMode:               true,
			DatabaseCache:           config.DaxxcoinDatabaseCache,
			NetworkId:               config.DaxxcoinNetworkID,
			GasPrice:                new(big.Int).Mul(big.NewInt(20), common.Shannon),
			GpoMinGasPrice:          new(big.Int).Mul(big.NewInt(20), common.Shannon),
			GpoMaxGasPrice:          new(big.Int).Mul(big.NewInt(500), common.Shannon),
			GpoFullBlockRatio:       80,
			GpobaseStepDown:         10,
			GpobaseStepUp:           100,
			GpobaseCorrectionFactor: 110,
		}
		if err := rawStack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, ethConf)
		}); err != nil {
			return nil, fmt.Errorf("daxxcoin init: %v", err)
		}
		// If netstats reporting is requested, do it
		if config.DaxxcoinNetStats != "" {
			if err := rawStack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
				var lesServ *les.LightDaxxcoin
				ctx.Service(&lesServ)

				return ethstats.New(config.DaxxcoinNetStats, nil, lesServ)
			}); err != nil {
				return nil, fmt.Errorf("netstats init: %v", err)
			}
		}
	}
	// Register the Whisper protocol if requested
	if config.WhisperEnabled {
		if err := rawStack.Register(func(*node.ServiceContext) (node.Service, error) { return whisperv2.New(), nil }); err != nil {
			return nil, fmt.Errorf("whisper init: %v", err)
		}
	}
	return &Node{rawStack}, nil
}

// Start creates a live P2P node and starts running it.
func (n *Node) Start() error {
	return n.node.Start()
}

// Stop terminates a running node along with all it's services. In the node was
// not started, an error is returned.
func (n *Node) Stop() error {
	return n.node.Stop()
}

// GetDaxxcoinClient retrieves a client to access the Daxxcoin subsystem.
func (n *Node) GetDaxxcoinClient() (client *DaxxcoinClient, _ error) {
	rpc, err := n.node.Attach()
	if err != nil {
		return nil, err
	}
	return &DaxxcoinClient{ethclient.NewClient(rpc)}, nil
}

// GetNodeInfo gathers and returns a collection of metadata known about the host.
func (n *Node) GetNodeInfo() *NodeInfo {
	return &NodeInfo{n.node.Server().NodeInfo()}
}

// GetPeersInfo returns an array of metadata objects describing connected peers.
func (n *Node) GetPeersInfo() *PeerInfos {
	return &PeerInfos{n.node.Server().PeersInfo()}
}
