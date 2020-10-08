# stagger

[![codecov](https://codecov.io/gh/deweppro/stagger/branch/master/graph/badge.svg)](https://codecov.io/gh/deweppro/stagger)
[![Release](https://img.shields.io/github/release/deweppro/stagger.svg?style=flat-square)](https://github.com/deweppro/stagger/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/deweppro/stagger)](https://goreportcard.com/report/github.com/deweppro/stagger)
[![Build Status](https://travis-ci.com/deweppro/stagger.svg?branch=master)](https://travis-ci.com/deweppro/stagger)

# Consul migration

Example:
```go
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

```

Migration file format:

```yaml
- key: test                   # name
  value: '{"hello":"world"}'  # value
  type: json                  # type (choice: base64, json)
```

### Add Keys

Default format: (without specifying the key type to disable value validation)
```yaml
- key: test  
  value: 123
```

JSON format: (with key value validation as json)
```yaml
- key: test  
  value: '[1,2,3,4]'
  type: json
```

Base64 format: (with key value validation as base64)
```yaml
- key: test  
  value: aGVsbG8gd29ybGQ=
  type: base64
```

### Delete Keys

For delete a key from the consul, remove the parameter with the key value in the migration file.
```yaml
- key: test
```