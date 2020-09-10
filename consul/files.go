package consul

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	ErrUnsupportedFormat = errors.New("unsupported format")
)

func DirDecode(path string) (KVList, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if IsAllowedExtension(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Strings(files)

	temp := make(map[string]KVItem)
	for _, f := range files {
		data, err := FileDecode(f)
		if err != nil {
			return nil, err
		}
		if err = data.Validate(); err != nil {
			return nil, err
		}
		for _, i := range data {
			temp[i.Key] = i
		}
	}

	kv := make(KVList, 0)
	for _, i := range temp {
		kv = append(kv, i)
	}

	return kv, nil
}

func FileDecode(filename string) (KVList, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}

	var kv KVList
	switch filepath.Ext(filename) {
	case ".json":
		if err := json.Unmarshal(body, &kv); err != nil {
			return nil, err
		}
	case ".yml", ".yaml":
		if err := yaml.Unmarshal(body, &kv); err != nil {
			return nil, err
		}
	default:
		return nil, ErrUnsupportedFormat
	}

	return kv, nil
}

func FileEncode(filename string, kv KVList) error {
	var (
		body []byte
		err  error
	)
	switch filepath.Ext(filename) {
	case ".json":
		if body, err = yaml.Marshal(kv); err != nil {
			return err
		}
	case ".yml", ".yaml":
		if body, err = yaml.Marshal(kv); err != nil {
			return err
		}
	default:
		return ErrUnsupportedFormat
	}

	return ioutil.WriteFile(filename, body, 0755)
}

func IsAllowedExtension(filename string) bool {
	switch filepath.Ext(filename) {
	case ".json", ".yml", ".yaml":
		return true
	}
	return false
}
