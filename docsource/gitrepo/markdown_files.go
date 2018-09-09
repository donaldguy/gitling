package gitrepo

import (
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func fileIsMarkdown(f *object.File) bool {
	return strings.HasSuffix(f.Name, ".md")
}

func headMarkdownFiles(repo *git.Repository) (results []*object.File, err error) {
	ref, err := repo.Head()
	if err != nil {
		return
	}

	c, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return
	}

	fs, err := c.Files()
	if err != nil {
		return
	}

	err = fs.ForEach(func(f *object.File) error {
		if fileIsMarkdown(f) {
			results = append(results, f)
		}
		return nil
	})

	return
}
