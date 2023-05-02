package filecopy

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const testPath = "../../test/"

func errorsEqual(err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}
	if err1 == nil || err2 == nil {
		return false
	}
	return err1.Error() == err2.Error()
}

func TestCopyFiles(t *testing.T) {
	type test struct {
		name      string
		sourceDir string
		fileMask  string
		wantError error
	}

	tests := []test{
		{
			name:      "file mask *",
			sourceDir: filepath.Join(testPath, "/source/"),
			fileMask:  "*",
			wantError: nil,
		},
		{
			name:      "file mask *.txt",
			sourceDir: filepath.Join(testPath, "/source/"),
			fileMask:  "*.txt",
			wantError: nil,
		},
		{
			name:      "file mask 2.*",
			sourceDir: filepath.Join(testPath, "/source/"),
			fileMask:  "2.*",
			wantError: nil,
		},
		{
			name:      "invalid path: E",
			sourceDir: "E",
			fileMask:  "*",
			wantError: fmt.Errorf("invalid path: %s", "E"),
		},
		{
			name:      "path is not a directory: /source/1.txt",
			sourceDir: filepath.Join(testPath, "/source/1.txt"),
			fileMask:  "*",
			wantError: fmt.Errorf("path is not a directory: %s", filepath.Join(testPath, "/source/1.txt")),
		},
	}

	var targetDir string

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.wantError == nil {
				var err error
				targetDir, err = os.MkdirTemp(testPath, "target")
				if err != nil {
					t.Errorf("invalid test case: %v", err)
					return
				}

				defer os.RemoveAll(targetDir)
			}

			gotError := CopyFiles(tt.sourceDir, tt.fileMask, targetDir)

			if !errorsEqual(gotError, tt.wantError) {
				t.Errorf("CopyFiles error got = %v, expected %v", gotError, tt.wantError)
				return
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
