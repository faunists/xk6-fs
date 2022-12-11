package fs

import (
	"os"
	"path/filepath"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/fs", new(FS))
}

// FS is the k6 extension
type FS struct{}

// WalkMatch walk through files and return all matches, otherwise an error.
// Usage: files, err := WalkMatch("/root/", "*.js")
func (*FS) WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
