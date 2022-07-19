package mapsupport

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		one    map[string]interface{}
		second map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "case 1",
			args: args{
				one: map[string]interface{}{
					"key-1": "1",
					"key-2": "2",
				},
				second: map[string]interface{}{
					"key-2": "minhnv",
				},
			},
			want: map[string]interface{}{
				"key-1": "1",
				"key-2": "minhnv",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.one, tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
