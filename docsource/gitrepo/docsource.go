package gitrepo

import (
	"fmt"

	"github.com/donaldguy/gitling/docsource"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type GitRepoDocSource struct {
	repo          *git.Repository
	markdownFiles map[string]*object.File
}

func New(path string) (*GitRepoDocSource, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	files, err := headMarkdownFiles(repo)
	if err != nil {
		return nil, err
	}

	fileMap := make(map[string]*object.File)
	for _, f := range files {
		fileMap[f.Name] = f
	}

	return &GitRepoDocSource{
		repo:          repo,
		markdownFiles: fileMap,
	}, err
}

var _ docsource.DocSource = (*GitRepoDocSource)(nil)

func (g *GitRepoDocSource) List() []string {
	names := make([]string, len(g.markdownFiles))
	i := 0
	for k := range g.markdownFiles {
		names[i] = k
		i++
	}
	return names
}

func (g *GitRepoDocSource) GetMarkdown(name string) (string, error) {
	f, ok := g.markdownFiles[name]
	if !ok {
		return "", fmt.Errorf("File %s not found in repo", name)
	}
	return f.Contents()
}

func (g *GitRepoDocSource) GetVersion(name string) (*docsource.Version, error) {
	changeCommit, err := fileLastChanged(g.repo, name)

	if err != nil {
		return nil, err
	}

	return &docsource.Version{
		ID:        changeCommit.Hash.String(),
		Timestamp: changeCommit.Committer.When,
	}, nil
}

func (g *GitRepoDocSource) GetLastAuthor(name string) (*docsource.Author, error) {
	changeCommit, err := fileLastChanged(g.repo, name)
	if err != nil {
		return nil, err
	}

	return &docsource.Author{
		Name:  changeCommit.Author.Name,
		Email: changeCommit.Author.Email,
	}, nil
}

func (g *GitRepoDocSource) GetFirstAuthor(name string) (*docsource.Author, error) {
	changeCommit, err := fileIntroduced(g.repo, name)
	if err != nil {
		return nil, err
	}

	return &docsource.Author{
		Name:  changeCommit.Author.Name,
		Email: changeCommit.Author.Email,
	}, nil
}
