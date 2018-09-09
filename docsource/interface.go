package docsource

import "time"

type DocSource interface {
	List() []string

	GetMarkdown(string) (string, error)
	GetVersion(string) (*Version, error)

	GetFirstAuthor(string) (*Author, error)
	GetLastAuthor(string) (*Author, error)
}

type Version struct {
	ID        string
	Timestamp time.Time
}

type Author struct {
	Name  string
	Email string
}
