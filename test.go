package main

import (
	"log"
	"net/http"
	"webgo/server"
)

type student struct {
	Name string
	Age  int
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()

	s := server.InitServer()

	s.LoadTemplate("test/templates/*")
	s.StaticResource("/static/css", "test/static")

	stu1 := &student{Name: "yanyibin", Age: 23}
	stu2 := &student{Name: "ty", Age: 23}

	s.Get("/student", func(c *server.Context) {
		c.HTML(http.StatusOK, "test.tmpl", server.Content{
			"title":    "yanyibin",
			"students": [2]*student{stu1, stu2},
		})
	})

	s.Get("/css", func(c *server.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	group := s.SetGroup("/pre")
	{

		group.Get("/:path/bbb/:yan", func(c *server.Context) {
			c.JSON(http.StatusOK, server.Content{
				"t1":       c.PathParams["yan"],
				"text":     c.PathParams["path"],
				"username": "yanyibin",
				"password": "yyb",
			})
		})
	}

	s.Get("/:path/bbb/:yan", func(c *server.Context) {
		c.JSON(http.StatusOK, server.Content{
			"t1":       c.PathParams["yan"],
			"text":     c.PathParams["path"],
			"username": "yanyibin",
			"password": "yyb",
		})
	})

	s.Run("localhost:9999")
}
