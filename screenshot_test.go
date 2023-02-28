package screenshot

import (
	"os"
	"testing"
)

func TestScreenshotURLs(t *testing.T) {
	tt := []struct {
		url     string
		wantErr bool
	}{
		{url: "https://www.golang.org", wantErr: false},
		{url: "bbc.com", wantErr: true},
		{url: "", wantErr: true},
	}
	for _, tt := range tt {
		_, err := screenshot(&tt.url)

		if err != nil && !tt.wantErr {
			t.Fatalf("Screenshot URL is invalid: '%s'", tt.url)
		}
	}
}

func TestWriteFile(t *testing.T) {
	tt := []struct {
		file    string
		wantErr bool
	}{
		{file: "filename.png", wantErr: false},
		{file: "", wantErr: true},
	}
	for _, tt := range tt {
		var buf []byte
		err := writeFile(&tt.file, buf)
		defer os.Remove(tt.file)

		if err != nil && !tt.wantErr {
			t.Fatalf("Failed to write file: '%s'", tt.file)
		}
	}
}
