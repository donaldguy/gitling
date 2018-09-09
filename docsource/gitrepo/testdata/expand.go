package testdata

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// git, somewhat understandably is pretty intransigent to check-in a .git
// folder into a git repo, so I had to tar the test data; and then I gziped
// it to save space. So here we have a helper that does tar -xzf in go
//
// a side-effect of this is that if we write tests that write to the repo
// later, this can act as a reset.
func ExpandTarball(name string, t *testing.T) (string, func()) {
	// We could do something under a tempdir instead, but since its .gitignored this seems more
	// convenient
	expectedTarget := filepath.Join("testdata", strings.TrimSuffix(name, filepath.Ext(name)))

	cleanupFunc := func() {
		err := os.RemoveAll(expectedTarget)
		if err != nil {
			t.Fatalf("Couldn't remove old testdata %s: %v", name, err)
		}
	}

	cleanupFunc()

	inf, err := os.Open(filepath.Join("testdata", name))

	if err != nil {
		t.Fatalf("Couldn't open testdata %s: %v", name, err)
	}
	defer inf.Close()

	gzf, err := gzip.NewReader(inf)
	if err != nil {
		t.Fatalf("Couldn't open testdata %s: %v", name, err)
	}
	tarf := tar.NewReader(gzf)
	if err != nil {
		t.Fatalf("Couldn't open testdata %s: %v", name, err)
	}
	for {
		header, err := tarf.Next()

		switch {
		case err == io.EOF:
			return expectedTarget, cleanupFunc
		case err != nil:
			t.Fatalf("Couldn't untar testdata %s: %v", name, err)
		case header == nil:
			continue
		}

		target := filepath.Join("testdata", header.Name)

		switch header.Typeflag {

		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					t.Fatalf("Couldn't untar testdata %s: %v", name, err)
				}
			}

		case tar.TypeReg:
			outf, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				t.Fatalf("Couldn't untar testdata %s: %v", name, err)
			}

			// copy over contents
			if _, err := io.Copy(outf, tarf); err != nil {
				t.Fatalf("Couldn't untar testdata %s: %v", name, err)
			}
			outf.Close()
		}
	}
}
