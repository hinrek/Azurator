package filereader

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Should pass with correct path and content",
			args: args{path: "../test/testdata/test_template.yml"},
			want: []byte("testyaml:"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
