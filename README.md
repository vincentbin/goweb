# webgo
A gin-like light web framework.
## Getting Start
支持动态、静态路由，分组路由配置，以及静态文件获取，模板渲染页面。
***
### 静态路由：
```golang
func main() {
    s := server.InitServer()
    s.Get("/level1/level2", func(c *server.Context) {
		  c.JSON(http.StatusOK, server.Content{
        "username": "yanyibin",
        "password": "yyb",
      })
    })
    s.Run("localhost:9999")
}
```
### 动态路由：
可由Context中PathParams中获得动态路由名
<br/>利用trie树进行实现
```golang
func main() {
    s := server.InitServer()
    s.Get("/level1/:v1", func(c *server.Context) {
      c.JSON(http.StatusOK, server.Content{
        "pathParam": c.PathParams["v1"]
        "username": "yanyibin",
        "password": "yyb",
      })
    })
    s.Run("localhost:9999")
}
```
### 分组路由：
```golang
func main() {
    s := server.InitServer()
    g := s.SetGroup("/group1")
    {
        g.Get("/level1/:v1", func(c *server.Context) {
          c.JSON(http.StatusOK, server.Content{
            "pathParam": c.PathParams["v1"]
            "username": "yanyibin",
            "password": "yyb",
          })
        })
        
        g.Get("/level2/:v2", func(c *server.Context) {
          c.JSON(http.StatusOK, server.Content{
            "pathParam": c.PathParams["v2"]
            "username": "yanyibin",
            "password": "yyb",
          })
        })
    }
    s.Run("localhost:9999")
}
```
### 静态文件访问、模板渲染：
通过localhost:9999/student 即可获取test.tmpl对应页面。
```golang
type student struct {
	Name string
	Age  int
}

func main() {
    s := server.InitServer()
    s.LoadTemplate("test/templates/*")
    s.StaticResource("/static/css", "test/static")
    
    s1 := &student{Name: "yanyibin", Age: 23}
	  s2 := &student{Name: "ty", Age: 23}
    
    s.Get("/student", func(c *server.Context) {
      c.HTML(http.StatusOK, "test.tmpl", server.Content{
        "title":    "yanyibin",
        "students": [2]*student{s1, s2},
      })
    })
    s.Run("localhost:9999")
}
```
***
## 目录结构
### 目录结构描述
```
.
├── README.md                   // readme
├── server                      
│   ├── context.go              // 请求上下文
│   ├── group.go                // 服务url前缀分组
│   ├── router.go               // 请求路由
│   ├── server.go               // 服务相关
|
├── util
|   |── string.go               // 字符串处理工具
|   |── trie.go                 // 实现动态路由 trie树
|
├── test                        // 静态文件测试用包
|   |── static                  // js & css
|   |── templates               // tmpl模板
|
├── test.go                     // 测试启动类
```
