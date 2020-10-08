/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package consul

import "testing"

func TestUnit_KVItem_Validate(t *testing.T) {
	type fields struct {
		Key   string
		Value string
		Type  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "bad type", fields: fields{Key: "a", Value: "{}", Type: "Json"}, wantErr: true},
		{name: "bad value", fields: fields{Key: "a", Value: "{}", Type: "base64"}, wantErr: true},
		{name: "bad key", fields: fields{Key: "", Value: "{}", Type: ""}, wantErr: true},
		{name: "good", fields: fields{Key: "a", Value: "{}", Type: "json"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kvi := KVItem{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
				Type:  tt.fields.Type,
			}
			if err := kvi.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestUnit_KVItem_DetectType(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		fields KVItem
	}{
		{name: "default", value: TypeDefault, fields: KVItem{Key: "", Value: "", Type: ""}},
		{name: "json", value: TypeJson, fields: KVItem{Key: "", Value: `{"a":"a"}`, Type: ""}},
		{name: "base64", value: TypeBase64, fields: KVItem{Key: "", Value: "aGVsbG8=", Type: ""}},
		{name: "bad_base64", value: TypeDefault, fields: KVItem{Key: "", Value: "?aGVsbG8=", Type: ""}},
		{name: "bad_json", value: TypeDefault, fields: KVItem{Key: "", Value: `["a":"a"]`, Type: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kvi := KVItem{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
				Type:  tt.fields.Type,
			}

			if ctype := kvi.DetectType(); ctype != tt.value {
				t.Errorf("DetectType() expect = %v, actual = %v", tt.value, ctype)
			}
		})
	}
}
