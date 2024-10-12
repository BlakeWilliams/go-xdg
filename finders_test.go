package xdg

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
)

func TestFindFile(t *testing.T) {
	testCases := []struct {
		desc       string
		env        map[string]string
		finder     func(tmpDir string) func() (string, error)
		error      string
		searchPath []string
	}{
		{
			desc:       "file exists",
			finder:     findTmpRoot,
			searchPath: []string{"foo.yml"},
		},
		{
			desc:       "file does not exist",
			finder:     findTmpRoot,
			searchPath: []string{"missing.yml"},
			error:      "no such file or directory",
		},
		{
			desc:       "finder fails",
			finder:     findWithError,
			searchPath: []string{"missing.yml"},
			error:      "oh no!",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tmpDir := t.TempDir()
			f, err := os.Create(path.Join(tmpDir, "foo.yml"))
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(f.Name())
			defer os.Remove(f.Name())
			f.Write([]byte(tC.desc))
			f.Close()

			fullPath, err := findFile(tC.finder(tmpDir), tC.searchPath...)
			if err != nil {
				if tC.error == "" {
					t.Fatalf("Expected no errors, got: %v", err)
				} else if !strings.HasSuffix(err.Error(), tC.error) {
					t.Errorf("Expected error %v, got %v", tC.error, err)
				}
				return
			}

			if !strings.HasSuffix(fullPath, path.Join(tC.searchPath...)) {
				t.Errorf("Expected %s, got %s", path.Join(tmpDir, "foo.yaml"), fullPath)
			}
		})
	}
}

func TestFindConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	resetEnv := stubEnv(map[string]string{
		"XDG_CONFIG_HOME": tmpDir,
	})
	defer resetEnv()

	f, err := os.Create(path.Join(tmpDir, "config.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	path, err := FindConfigFile("config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if path != f.Name() {
		t.Errorf("Expected %s, got %s", f.Name(), path)
	}
}

func findTmpRoot(tmpDir string) func() (string, error) {
	return func() (string, error) {
		return tmpDir, nil
	}
}
func findWithError(tmpDir string) func() (string, error) {
	return func() (string, error) {
		return "", fmt.Errorf("oh no!")
	}
}
