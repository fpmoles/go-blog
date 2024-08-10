package go_blog

import (
	"reflect"
	"testing"
)

func Test_loadConfig(t *testing.T) {
	tests := []struct {
		name       string
		configFile string
		want       *Config
		wantErr    bool
	}{
		{
			name:       "Happy Path",
			configFile: "./.github/testdata/test-config.yaml",
			want: &Config{
				Name: "foo",
				Site: "http://foo.com",
			},
			wantErr: false,
		},
		{
			name:       "Bad File",
			configFile: "./.github/testdata/test-config-bad.yaml",
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "No File",
			configFile: "foobar",
			want:       nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadConfig(tt.configFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
