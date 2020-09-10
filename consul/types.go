package consul

import (
	"encoding/base64"
	"encoding/json"

	"github.com/pkg/errors"
)

const (
	TypeDefault = ""
	TypeJson    = "json"
	TypeBase64  = "base64"
)

var (
	ErrUndefinedType = errors.New("undefined key type")
)

type (
	KVList []KVItem

	KVItem struct {
		Key   string `yaml:"key" json:"key"`
		Value string `yaml:"value,omitempty" json:"value,omitempty"`
		Type  string `yaml:"type,omitempty" json:"type,omitempty"`
	}
)

func (kvl KVList) Validate() error {
	var (
		err error
		er  error
	)
	for _, i := range kvl {
		if er = i.Validate(); er != nil {
			err = errors.WithMessagef(err, "key %s error %s", i.Key, er.Error())
		}
	}
	return err
}

func (kvi KVItem) ValueBytes() ([]byte, error) {
	if kvi.Type == TypeBase64 {
		return base64.StdEncoding.DecodeString(kvi.Value)
	}
	return []byte(kvi.Value), nil
}

func (kvi KVItem) Validate() error {
	if len(kvi.Key) == 0 {
		return errors.New("empty key")
	}

	switch kvi.Type {
	case TypeDefault:
	case TypeJson:
		var t json.RawMessage
		if err := json.Unmarshal([]byte(kvi.Value), &t); err != nil {
			return err
		}
	case TypeBase64:
		if _, err := base64.StdEncoding.DecodeString(kvi.Value); err != nil {
			return err
		}
	default:
		return errors.WithMessagef(ErrUndefinedType, "key %s type %s", kvi.Key, kvi.Type)
	}

	return nil
}
