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

package main

import (
	"log"

	"github.com/daxxcoin/daxxcore/swarm/storage"
	"gopkg.in/urfave/cli.v1"
)

func cleandb(ctx *cli.Context) {
	args := ctx.Args()
	if len(args) != 1 {
		log.Fatal("need path to chunks database as the first and only argument")
	}

	chunkDbPath := args[0]
	hash := storage.MakeHashFunc("SHA3")
	dbStore, err := storage.NewDbStore(chunkDbPath, hash, 10000000, 0)
	if err != nil {
		log.Fatalf("cannot initialise dbstore: %v", err)
	}
	dbStore.Cleanup()
}
