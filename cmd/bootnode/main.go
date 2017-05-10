// Copyright 2015 The daxxcoreAuthors
// This file is part of daxxCore.
//
// daxxcoreis free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// daxxcoreis distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with daxxCore. If not, see <http://www.gnu.org/licenses/>.

// bootnode runs a bootstrap node for the Daxxcoin Discovery Protocol.
package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"os"

	"github.com/daxxcoin/daxxcore/cmd/utils"
	"github.com/daxxcoin/daxxcore/crypto"
	"github.com/daxxcoin/daxxcore/logger/glog"
	"github.com/daxxcoin/daxxcore/p2p/discover"
	"github.com/daxxcoin/daxxcore/p2p/discv5"
	"github.com/daxxcoin/daxxcore/p2p/nat"
	"github.com/daxxcoin/daxxcore/p2p/netutil"
)

func main() {
	var (
		listenAddr  = flag.String("addr", ":30301", "listen address")
		genKey      = flag.String("genkey", "", "generate a node key")
		writeAddr   = flag.Bool("writeaddress", false, "write out the node's pubkey hash and quit")
		nodeKeyFile = flag.String("nodekey", "", "private key filename")
		nodeKeyHex  = flag.String("nodekeyhex", "", "private key as hex (for testing)")
		natdesc     = flag.String("nat", "none", "port mapping mechanism (any|none|upnp|pmp|extip:<IP>)")
		netrestrict = flag.String("netrestrict", "", "restrict network communication to the given IP networks (CIDR masks)")
		runv5       = flag.Bool("v5", false, "run a v5 topic discovery bootnode")

		nodeKey *ecdsa.PrivateKey
		err     error
	)
	flag.Var(glog.GetVerbosity(), "verbosity", "log verbosity (0-9)")
	flag.Var(glog.GetVModule(), "vmodule", "log verbosity pattern")
	glog.SetToStderr(true)
	flag.Parse()

	natm, err := nat.Parse(*natdesc)
	if err != nil {
		utils.Fatalf("-nat: %v", err)
	}
	switch {
	case *genKey != "":
		nodeKey, err = crypto.GenerateKey()
		if err != nil {
			utils.Fatalf("could not generate key: %v", err)
		}
		if err = crypto.SaveECDSA(*genKey, nodeKey); err != nil {
			utils.Fatalf("%v", err)
		}
	case *nodeKeyFile == "" && *nodeKeyHex == "":
		utils.Fatalf("Use -nodekey or -nodekeyhex to specify a private key")
	case *nodeKeyFile != "" && *nodeKeyHex != "":
		utils.Fatalf("Options -nodekey and -nodekeyhex are mutually exclusive")
	case *nodeKeyFile != "":
		if nodeKey, err = crypto.LoadECDSA(*nodeKeyFile); err != nil {
			utils.Fatalf("-nodekey: %v", err)
		}
	case *nodeKeyHex != "":
		if nodeKey, err = crypto.HexToECDSA(*nodeKeyHex); err != nil {
			utils.Fatalf("-nodekeyhex: %v", err)
		}
	}

	if *writeAddr {
		fmt.Printf("%v\n", discover.PubkeyID(&nodeKey.PublicKey))
		os.Exit(0)
	}

	var restrictList *netutil.Netlist
	if *netrestrict != "" {
		restrictList, err = netutil.ParseNetlist(*netrestrict)
		if err != nil {
			utils.Fatalf("-netrestrict: %v", err)
		}
	}

	if *runv5 {
		if _, err := discv5.ListenUDP(nodeKey, *listenAddr, natm, "", restrictList); err != nil {
			utils.Fatalf("%v", err)
		}
	} else {
		if _, err := discover.ListenUDP(nodeKey, *listenAddr, natm, "", restrictList); err != nil {
			utils.Fatalf("%v", err)
		}
	}

	select {}
}