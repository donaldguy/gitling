package gitrepo

import (
	"testing"

	"github.com/deckarep/golang-set"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/filemode"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"github.com/donaldguy/gitling/docsource/gitrepo/testdata"
)

var nilBlob = &object.Blob{}

func TestFileIsMarkdown(t *testing.T) {
	aMarkdownFile := object.NewFile("example.md", filemode.Regular, nilBlob)
	aNestedMarkdownFile := object.NewFile("sub/example.md", filemode.Regular, nilBlob)
	notAMarkdownFile := object.NewFile("example.rb", filemode.Regular, nilBlob)
	notANestedMarkdownFile := object.NewFile("sub/example.rb", filemode.Regular, nilBlob)

	if !fileIsMarkdown(aMarkdownFile) {
		t.Errorf("%s should be treated as markdown", aMarkdownFile.Name)
	}

	if !fileIsMarkdown(aNestedMarkdownFile) {
		t.Errorf("%s should be treated as markdown", aNestedMarkdownFile.Name)
	}

	if fileIsMarkdown(notAMarkdownFile) {
		t.Errorf("%s should NOT be treated as markdown", notAMarkdownFile.Name)
	}

	if fileIsMarkdown(notANestedMarkdownFile) {
		t.Errorf("%s should NOT be treated as markdown", notAMarkdownFile.Name)
	}
}

func TestHeadMarkdownFiles(t *testing.T) {
	repoDir, cleanup := testdata.ExpandTarball("repo_a.tgz", t)
	defer cleanup()

	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		t.Fatalf("Couldn't open repo: %v", err)
	}

	markdownFiles, err := headMarkdownFiles(repo)
	if err != nil {
		t.Errorf("Failed to parse markdown files out: %v", err)
	}

	outputSet := mapset.NewSet()
	for _, f := range markdownFiles {
		outputSet.Add(f.Name)
	}

	repoAMarkdownFiles := testdata.RepoAMarkdownFiles()
	if !outputSet.Equal(repoAMarkdownFiles) {
		t.Errorf("Expected %+v but got %+v", repoAMarkdownFiles, outputSet)
	}
}
