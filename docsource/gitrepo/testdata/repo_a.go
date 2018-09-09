package testdata

import (
	mapset "github.com/deckarep/golang-set"
)

func RepoAMarkdownFiles() mapset.Set {
	filenames := mapset.NewSet()

	filenames.Add("EXAMPLE.md")
	filenames.Add("subdir/nested.md")
	filenames.Add("subdir2/sister.md")

	return filenames
}

func RepoAFileVersions() map[string]string {
	return map[string]string{
		"EXAMPLE.md":                    "c60c80802f58b7cd78be11d3efc565a019733165",
		"subdir2/brother.js":            "0b2a52d54bd51746973e9ce617d43c61b3838e75",
		"subdir2/sister.md":             "18bb07f355883c389f11c8e906ed7c86cfe09cd2",
		"subdir/nested.md":              "18bb07f355883c389f11c8e906ed7c86cfe09cd2",
		"no_md_here/uninteresting.html": "18bb07f355883c389f11c8e906ed7c86fe09cd2",
		"not_markdown.sh":               "c60c80802f58b7cd78be11d3efc565a019733165",
	}
}
