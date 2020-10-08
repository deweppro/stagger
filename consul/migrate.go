/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
)

const (
	chunkSize = 64
)

type (
	ConsulMigrate struct {
		cli *api.Client
	}
)

func CreateConsulMigrate(addr string) (*ConsulMigrate, error) {
	cli, err := api.NewClient(&api.Config{Address: addr})
	if err != nil {
		return nil, err
	}

	return NewConsulMigrate(cli), nil
}

func NewConsulMigrate(cli *api.Client) *ConsulMigrate {
	return &ConsulMigrate{
		cli: cli,
	}
}

func (m *ConsulMigrate) DumpToFile(filename string) error {
	if !IsAllowedExtension(filename) {
		return ErrUnsupportedFormat
	}
	kv, err := m.Dump()
	if err != nil {
		return err
	}
	return FileEncode(filename, kv)
}

func (m *ConsulMigrate) Dump() (KVList, error) {
	kv, _, err := m.cli.KV().List("", nil)
	if err != nil {
		return nil, errors.Wrap(err, "get keys from consul")
	}
	result := make(KVList, 0)
	for _, i := range kv {
		nkv := KVItem{
			Key:   i.Key,
			Value: string(i.Value),
			Type:  "",
		}
		nkv.Type = nkv.DetectType()
		result = append(result, nkv)
	}
	return result, nil
}

func (m *ConsulMigrate) MigrateFromFile(filename string) error {
	kv, err := FileDecode(filename)
	if err != nil {
		return err
	}

	return m.Migrate(kv)
}

func (m *ConsulMigrate) MigrateFromDir(path string) error {
	kv, err := DirDecode(path)
	if err != nil {
		return err
	}

	return m.Migrate(kv)
}

func (m *ConsulMigrate) Migrate(kv KVList) error {
	if err := kv.Validate(); err != nil {
		return err
	}

	result := make(api.KVTxnOps, 0)
	for _, i := range kv {
		if len(i.Value) == 0 {
			result = append(result, &api.KVTxnOp{
				Verb: api.KVDelete,
				Key:  i.Key,
			})
		} else {
			v, err := i.ValueBytes()
			if err != nil {
				return err
			}
			result = append(result, &api.KVTxnOp{
				Verb:  api.KVSet,
				Key:   i.Key,
				Value: v,
			})
		}
	}

	var chunk api.KVTxnOps
	chunks := make([]api.KVTxnOps, 0)

	for len(result) >= chunkSize {
		chunk, result = result[:chunkSize], result[chunkSize:]
		chunks = append(chunks, chunk)
	}

	if len(result) > 0 {
		chunks = append(chunks, result[:len(result)])
	}

	for _, item := range chunks {
		ok, _, _, err := m.cli.KV().Txn(item, nil)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("can`t save data to consul")
		}
	}

	return nil
}
