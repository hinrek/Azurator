package utils

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Should pass with correct path and content",
			args: args{
				path: "../test/testdata/test_filereader.yml",
			},
			want:    []byte("this is a test file for utils"),
			wantErr: false,
		},
		{
			name:    "Should error with incorrect path",
			args:    args{path: "wrongFile.txt"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
