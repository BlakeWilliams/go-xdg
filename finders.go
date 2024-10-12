package xdg

import (
	"fmt"
	"os"
	"path"
)

// ErrNoFile represents when a file cannot be found when searching for a
// specific file
type ErrNoFile struct {
	rootError error
	root      string
	file      string
}

func (e ErrNoFile) Error() string {
	if e.rootError == nil {
		return fmt.Sprintf("could not find file %s in %s", e.file, e.root)
	} else if e.root == "" && e.rootError != nil {
		return fmt.Sprintf("could not find file %ss: %s", e.file, e.rootError)
	} else {
		return fmt.Sprintf("could not find file %s in %s: %s", e.file, e.root, e.rootError)
	}
}

func (e ErrNoFile) Unwrap() error {
	return e.rootError
}

// FindConfigFile searches for a file in the user's configuration directory. If
// the file is not found ErrNoFile is returned.
func FindConfigFile(pathParts ...string) (string, error) {
	return findFile(ConfigHome, pathParts...)
}

// FindConfigFile searches for a file in the user's data directory. If the file
// is not found ErrNoFile is returned.
func FindDataFile(pathParts ...string) (string, error) {
	return findFile(DataHome, pathParts...)
}

// FindCacheFile searches for a file in the user's cache directory. If the file
// is not found ErrNoFile is returned.
func FindCacheFile(pathParts ...string) (string, error) {
	return findFile(CacheHome, pathParts...)
}

// FindStateFile searches for a file in the user's state directory. If the file
// is not found ErrNoFile is returned.
func FindStateFile(pathParts ...string) (string, error) {
	return findFile(CacheHome, pathParts...)
}

func findFile(findRoot func() (string, error), pathParts ...string) (string, error) {
	configPath, err := findRoot()

	if err != nil {
		return "", ErrNoFile{
			rootError: err,
			file:      path.Join(pathParts...),
		}
	}

	fullPath := path.Join(configPath, path.Join(pathParts...))
	_, err = os.Stat(fullPath)
	if err != nil {
		return "", ErrNoFile{
			rootError: err,
			file:      path.Join(pathParts...),
			root:      configPath,
		}
	}

	return fullPath, nil
}
