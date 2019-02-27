package configuration

import (
	"reflect"
	"testing"
)

func TestConf_Configuration(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		c    *Conf
		args args
		want *Conf
	}{
		{
			name: "Should read and match values from test_template.yml",
			c: &Conf{
				struct {
					Name                string `yaml:"name"`
					PersonalAccessToken string `yaml:"personalAccessToken"`
					APIVersion          string `yaml:"apiVersion"`
				}{Name: "", PersonalAccessToken: "", APIVersion: ""},
				struct {
					Name                string `yaml:"name"`
					PersonalAccessToken string `yaml:"personalAccessToken"`
					APIVersion          string `yaml:"apiVersion"`
				}{Name: "", PersonalAccessToken: "", APIVersion: ""},
			},
			args: args{
				path: "../../test/testdata/test_template.yml",
			},
			want: &Conf{
				struct {
					Name                string `yaml:"name"`
					PersonalAccessToken string `yaml:"personalAccessToken"`
					APIVersion          string `yaml:"apiVersion"`
				}{Name: "test1", PersonalAccessToken: "asdasd", APIVersion: "5.0"},
				struct {
					Name                string `yaml:"name"`
					PersonalAccessToken string `yaml:"personalAccessToken"`
					APIVersion          string `yaml:"apiVersion"`
				}{Name: "test2", PersonalAccessToken: "dsadsa", APIVersion: "5.0"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Configuration(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conf.Configuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
