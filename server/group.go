package server

import (
	"net/http"
	"path"
)

type Group struct {
	server *Server
	urlPrefix string
}

func (g *Group) SetGroup(prefix string) *Group {
	s := g.server
	newGroup := &Group{
		server: s,
		urlPrefix: prefix,
	}
	return newGroup
}

func (g *Group) addRouter(t string, postUrl string, handler Handler) {
	url := g.urlPrefix + postUrl
	g.server.router.addRouter(t, url, handler)
}

func (g *Group) Get(url string, handler Handler) {
	g.addRouter("GET", url, handler)
}

func (g *Group) Post(url string, handler Handler) {
	g.addRouter("Post", url, handler)
}

func (g *Group) StaticResource(relativePath string, root string) {
	absolutePath := path.Join(g.urlPrefix, relativePath)
	fs := http.Dir(root)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	handler := func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
	url := path.Join(relativePath, "/:filepath")
	g.Get(url, handler)
}
