package server

import (
	"net/http"
	"webgo/util"
)

type Router struct {
	tree map[string]*util.Node
	routerMap map[string]Handler
}

func InitRouter() *Router {
	return &Router{
		tree: make(map[string]*util.Node),
		routerMap: make(map[string]Handler),
	}
}

func (r *Router) addRouter(t string, url string, handler Handler) {
	parts := util.ParseUrl(url)
	key := t + "-" + url
	_, ok := r.tree[t]
	if !ok {
		r.tree[t] = &util.Node{
			Children: make(map[string]*util.Node),
		}
	}
	r.tree[t].Insert(url, parts, 0)
	r.routerMap[key] = handler
}

func (r *Router) getRouter(t string, url string) (string, map[string]string) {
	inputParts := util.ParseUrl(url)
	pathParams := make(map[string]string)
	root, ok := r.tree[t]
	if !ok {
		return "", nil
	}
	methodUrl := root.Search(inputParts, 0)
	if methodUrl == "" {
		return "", nil
	}
	parts := util.ParseUrl(methodUrl)
	for i, v := range parts {
		if v[0] == ':' {
			pathParams[v[1 : ]] = inputParts[i]
		}
	}
	return methodUrl, pathParams
}

func (r *Router) handle(c *Context) {
	url, pathParams := r.getRouter(c.Method, c.Path)
	if url == "" {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s.\n", c.Path)
		return
	}
	c.PathParams = pathParams
	key := c.Method + "-" + url
	r.routerMap[key](c)
}
