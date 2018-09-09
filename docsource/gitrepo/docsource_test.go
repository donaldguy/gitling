package gitrepo

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/deckarep/golang-set"

	"github.com/donaldguy/gitling/docsource/gitrepo/testdata"
)

func TestList(t *testing.T) {
	repoDir, cleanup := testdata.ExpandTarball("repo_a.tgz", t)
	defer cleanup()

	docsource, err := New(repoDir)
	if err != nil {
		t.Fatalf("Couldn't open docsource: %v", err)
	}

	listedDocs := docsource.List()

	outputSet := mapset.NewSet()
	for _, f := range listedDocs {
		outputSet.Add(f)
	}

	repoAMarkdownFiles := testdata.RepoAMarkdownFiles()
	if !outputSet.Equal(repoAMarkdownFiles) {
		t.Errorf("Expected %+v but got %+v", repoAMarkdownFiles, outputSet)
	}
}

// we could hardcode (or go-bindata) the file content, but I am gonna content
// with comparing whats coming out of the gitrepo with whats in the working copy
// since I know they were the same when I tarred them
func TestGetMarkdown(t *testing.T) {
	repoDir, cleanup := testdata.ExpandTarball("repo_a.tgz", t)
	defer cleanup()
	it := testdata.RepoAMarkdownFiles().Iterator()

	docsource, err := New(repoDir)
	if err != nil {
		t.Fatalf("Couldn't open docsource: %v", err)
	}

	for f := range it.C {
		fileName, ok := f.(string)
		if !ok {
			t.Fatalf("Unexpected Type assertion Failur!")
		}

		fromDisk, err := ioutil.ReadFile(filepath.Join(repoDir, fileName))
		if err != nil {
			t.Fatalf("Failed to read file '%s' from disk: %v", fileName, err)
		}

		viaGit, err := docsource.GetMarkdown(fileName)
		if err != nil {
			t.Fatalf("Failed to read file '%s' via git: %v", fileName, err)
		}

		if viaGit != string(fromDisk) {
			t.Errorf("%s differed from git and disk", fileName)
		}
		t.Logf("Checked %s", fileName)
	}
}
func TestGetVersion(t *testing.T) {
	repoDir, cleanup := testdata.ExpandTarball("repo_a.tgz", t)
	defer cleanup()
	it := testdata.RepoAMarkdownFiles().Iterator()
	repoAFileMap := testdata.RepoAFileVersions()

	docsource, err := New(repoDir)
	if err != nil {
		t.Fatalf("Couldn't open docsource: %v", err)
	}

	for f := range it.C {
		fileName, ok := f.(string)
		if !ok {
			t.Fatalf("Unexpected Type assertion Failur!")
		}

		expected, ok := repoAFileMap[fileName]
		if !ok {
			t.Fatalf("File map and version map are out of sync!")
		}

		actual, err := docsource.GetVersion(fileName)
		if err != nil {
			t.Fatalf("Failed to find version of '%s' via git: %v", fileName, err)
		}

		if expected != actual.ID {
			t.Errorf("%s version: expected '%s', got '%s'", fileName, expected, actual.ID)
		}
		t.Logf("Checked %s", fileName)
	}
}
