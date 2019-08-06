package main

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_parsePackageResolutionFile(t *testing.T) {
	var analyzerPath url.URL
	analyzerPath.Path = "/Users/vagrant/.pub-cache/hosted/pub.dartlang.org/analyzer-0.36.4/lib/"
	analyzerPath.Scheme = "file"

	var relPath url.URL
	relPath.Path = "../../.pub-cache/hosted/pub.dartlang.org/analyzer-0.36.4/lib/"

	tests := []struct {
		name     string
		contents string
		want     map[string]url.URL
		wantErr  bool
	}{
		{
			name: "empty file",
			contents: `# Generated by pub on 2019-08-05 14:50:08.261783.

# Other comment`,
			want:    map[string]url.URL{},
			wantErr: false,
		},
		{
			name: "",
			contents: `# Generated by pub on 2019-08-05 14:50:08.261783.
analyzer:file:///Users/vagrant/.pub-cache/hosted/pub.dartlang.org/analyzer-0.36.4/lib/`,
			want: map[string]url.URL{
				"analyzer": analyzerPath,
			},
			wantErr: false,
		},
		{
			name: "",
			contents: `# Generated by pub on 2019-08-05 14:50:08.261783.
analyzer:../../.pub-cache/hosted/pub.dartlang.org/analyzer-0.36.4/lib/`,
			want: map[string]url.URL{
				"analyzer": relPath,
			},
			wantErr: false,
		},
		{
			name: "",
			contents: `# Generated by pub on 2019-08-05 14:50:08.261783.
analyzer::invalid/ss`,
			want:    map[string]url.URL{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsePackageResolutionFile(tt.contents)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePackageResolutionFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePackageResolutionFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
