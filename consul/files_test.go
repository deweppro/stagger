/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package consul

import (
	"io/ioutil"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnit_DirDecode(t *testing.T) {
	path, err := ioutil.TempDir("/tmp", "TestDirDecode_*")
	require.NoError(t, err)
	dataYaml := `
- key: aaa
  value: aaa
`
	err = ioutil.WriteFile(path+"/demo1.yaml", []byte(dataYaml), 0755)
	require.NoError(t, err)

	dataJson := `
[{"key":"bbb", "value":"bbb"}]
`
	err = ioutil.WriteFile(path+"/demo1.json", []byte(dataJson), 0755)
	require.NoError(t, err)

	actual, err := DirDecode("/tmp")
	require.NoError(t, err)

	required := KVList{
		{Key: "aaa", Value: "aaa", Type: ""},
		{Key: "bbb", Value: "bbb", Type: ""},
	}

	sort.Slice(actual, func(i, j int) bool {
		return actual[i].Key < actual[j].Key
	})

	require.Equal(t, required, actual, "required %v actual %v", required, actual)
}

func TestUnit_FileDecode(t *testing.T) {
	path, err := ioutil.TempDir("/tmp", "TestDirDecode_*")
	require.NoError(t, err)
	dataYaml := `
- key: aaa
  value: aaa
`
	err = ioutil.WriteFile(path+"/demo1.yaml", []byte(dataYaml), 0755)
	require.NoError(t, err)

	actual, err := FileDecode(path + "/demo1.yaml")
	require.NoError(t, err)

	required := KVList{
		{Key: "aaa", Value: "aaa", Type: ""},
	}

	require.Equal(t, required, actual, "required %v actual %v", required, actual)
}

func TestUnit_FileEncode(t *testing.T) {
	path, err := ioutil.TempDir("/tmp", "TestDirDecode_*")
	require.NoError(t, err)
	required := `- key: aaa
  value: aaa
`

	data := KVList{
		{Key: "aaa", Value: "aaa", Type: ""},
	}
	require.NoError(t, FileEncode(path+"/demo1.yaml", data))

	actual, err := ioutil.ReadFile(path + "/demo1.yaml")
	require.NoError(t, err)

	require.Equal(t, required, string(actual), "required %v actual %v", required, string(actual))
}

func TestUnit_IsAllowedExtension(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{name: ".js", filename: "aaa.js", want: false},
		{name: ".json", filename: "aaa.json", want: true},
		{name: ".yaml", filename: "aaa.yaml", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllowedExtension(tt.filename); got != tt.want {
				t.Errorf("IsAllowedExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
