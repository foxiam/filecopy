package filecopy_test

import (
	"filecopy/pkg/filecopy"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFiles(t *testing.T) {
	type test struct {
		name      string
		sourceDir string
		fileMask  string
		targetDir string
		wantError error
	}

	tests := []test{
		{
			sourceDir: "./source/",
			name:      "file mask *",
			targetDir: "./tetst1/",
			wantError: nil,
		},
		{
			name:      "file mask *.txt",
			sourceDir: "./source/",
			fileMask:  "*.txt",
			targetDir: "./tetst2/",
			wantError: nil,
		},
		{
			name:      "file mask 2.*",
			sourceDir: "./source/",
			fileMask:  "2.*",
			targetDir: "./tetst3/",
			wantError: nil,
		},
		{
			name:      "invalid path: E",
			sourceDir: "E",
			fileMask:  "*",
			targetDir: "",
			wantError: fmt.Errorf("invalid path: %s", "E"),
		},
		{
			name:      "path is not a directory: E:/source/1.txt",
			sourceDir: "./source/1.txt",
			fileMask:  "*",
			targetDir: "",
			wantError: fmt.Errorf("path is not a directory: %s", "./source/1.txt"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.wantError == nil {
				err := os.MkdirAll(tt.targetDir, os.ModePerm)
				if err != nil {
					t.Errorf("invalid test case %v", err)
				}

				defer os.RemoveAll(filepath.Dir(tt.targetDir))
			}

			gotError := filecopy.CopyFiles(tt.sourceDir, tt.fileMask, tt.targetDir)
			if gotError != nil && gotError.Error() != tt.wantError.Error() {
				t.Errorf("CopyFiles error got = %v, expected %v", gotError, tt.wantError)
			}

			if gotError == nil {
				err := filepath.Walk(tt.sourceDir, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					matched, err := filepath.Match(tt.fileMask, info.Name())
					if err != nil {
						return err
					}

					if !info.IsDir() && matched {
						_, err := os.Stat(path)
						if err != nil {
							t.Errorf("file %v not found", path)
						}
					}
					return nil
				})

				if err != nil {
					t.Errorf("error = %v", err)
				}
			}
		})
	}
}
