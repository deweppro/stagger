package consul

import "testing"

func TestKVItem_Validate(t *testing.T) {
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
