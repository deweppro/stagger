/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"

	"github.com/deweppro/stagger/consul"
)

func main() {

	// Init migration client
	cli, err := consul.CreateConsulMigrate("consul:8500")
	if err != nil {
		panic(err)
	}

	// Get all existing keys from consul
	kvs, err := cli.Dump()
	if err != nil {
		panic(err)
	}
	for _, kv := range kvs {
		fmt.Println(kv.Key, kv.Type, kv.Value)
	}

	// Save all existing keys from consul to file
	err = cli.DumpToFile("/tml/dump.yaml")
	if err != nil {
		panic(err)
	}

	// Save keys to consul
	err = cli.Migrate(kvs)
	if err != nil {
		panic(err)
	}

	// Migration of keys from a file in the consul.
	err = cli.MigrateFromFile("/tmp/dump.yaml")
	if err != nil {
		panic(err)
	}

	// Scan directories, merge files with sort by name,
	// and the migration of keys from files in the consul.
	err = cli.MigrateFromDir("/tmp")
	if err != nil {
		panic(err)
	}

}
