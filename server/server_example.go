package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/donaldguy/gitling/docsource/gitrepo"
	"github.com/gin-gonic/gin"
)

func useDevProxy() bool {
	return true
}

func devProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		target, err := url.Parse("http://localhost:1234/")
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func repos(repoPaths map[string]string) map[string]*gitrepo.GitRepoDocSource {
	repos := make(map[string]*gitrepo.GitRepoDocSource, len(repoPaths))

	for k, v := range repoPaths {
		r, err := gitrepo.New(v)
		if err != nil {
			panic(err)
		}
		repos[k] = r
	}

	return repos
}

func checkRepo(repos map[string]*gitrepo.GitRepoDocSource, c *gin.Context) *gitrepo.GitRepoDocSource {
	repo, ok := repos[c.Param("repo")]
	if !ok {
		c.String(http.StatusNotFound, "repo %s not found", c.Param("repo"))
		return nil
	}
	return repo
}

func main() {
	repoPaths := map[string]string{
		"gitling": os.Args[1],
	}
	repos := repos(repoPaths)

	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	r := gin.Default()

	r.GET("/api/repos", func(c *gin.Context) {
		repoNames := make([]string, len(repos))
		i := 0
		for k := range repos {
			repoNames[i] = k
			i++
		}
		c.JSON(http.StatusOK, repoNames)
	})

	r.GET("/api/repos/:repo", func(c *gin.Context) {
		if repo := checkRepo(repos, c); repo != nil {
			c.JSON(http.StatusOK, map[string][]string{"files": repo.List()})
		}
	})

	r.GET("/api/repos/:repo/*file", func(c *gin.Context) {
		repo := checkRepo(repos, c)
		if repo == nil {
			return
		}

		file := strings.TrimPrefix(c.Param("file"), "/")

		if strings.HasSuffix(file, "/version") {
			v, _ := repo.GetVersion(strings.TrimSuffix(file, "/version"))
			c.JSON(http.StatusOK, v)
			return
		}

		if strings.HasSuffix(file, "/author") {
			a, _ := repo.GetLastAuthor(strings.TrimSuffix(file, "/author"))
			c.JSON(http.StatusOK, a)
			return
		}

		if strings.HasSuffix(file, "/op") {
			a, _ := repo.GetFirstAuthor(strings.TrimSuffix(file, "/op"))
			c.JSON(http.StatusOK, a)
			return
		}

		s, err := repo.GetMarkdown(strings.TrimPrefix(c.Param("file"), "/"))
		if err != nil {
			c.Error(err)
		} else {
			c.String(http.StatusOK, "%s", s)
		}
	})

	if useDevProxy() {
		r.NoRoute(devProxy())
	}

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":8080", "./server/testdata/server.pem", "./server/testdata/server.key")
}
