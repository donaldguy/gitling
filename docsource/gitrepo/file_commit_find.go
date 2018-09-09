package gitrepo

import (
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
)

// returns the commit which most recently modified the file
func fileLastChanged(repo *git.Repository, name string) (*object.Commit, error) {
	return fileCommitFind(repo, name, false)
}

// returns the commit that first introduced the file
func fileIntroduced(repo *git.Repository, name string) (*object.Commit, error) {
	return fileCommitFind(repo, name, true)
}

// simulation of git log -- $name, but probably even a bit more hueristic, to
// finds the commit that introduced (or if addedOnly == false, lastModified) the file
func fileCommitFind(repo *git.Repository, name string, addedOnly bool) (*object.Commit, error) {
	var blobHashAtHead plumbing.Hash
	var prevCommit *object.Commit

	headRef, err := repo.Head()
	if err != nil {
		return nil, err
	}

	headC, err := repo.CommitObject(headRef.Hash())
	if err != nil {
		return nil, err
	}

	f, err := headC.File(name)
	if err != nil {
		return nil, err
	}

	blobHashAtHead = f.Hash
	prevCommit = headC

	commits, err := repo.Log(&git.LogOptions{})
	if err != nil {
		return nil, err
	}

	err = commits.ForEach(func(c *object.Commit) error {
		if c.Hash == prevCommit.Hash { // i.e. this is HEAD
			return nil //continue
		}

		f, err := c.File(name)
		if err == object.ErrFileNotFound { //it was created in commit checked last
			return storer.ErrStop
		}
		if err != nil {
			return err
		}

		if !addedOnly && f.Hash != blobHashAtHead { //it was changed in commit checked last
			return storer.ErrStop
		}

		prevCommit = c
		return nil
	})
	if err != nil {
		return nil, err
	}

	return prevCommit, nil
}
