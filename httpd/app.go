package main

import (
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"joshjwelsh.github.io/httpd/handler"
)

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	fmt.Println("TEMPLATES DIR: ", templatesDir)
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	fmt.Println("Layouts: ", layouts)
	if err != nil {
		fmt.Println("Houston theres a problem 1")
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		fmt.Println("Houston theres a problem 2")

		panic(err.Error())
	}
	var f []string
	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
		f = files
	}
	fmt.Println("FILES: ", f)
	return r
}

// func createMyRender() multitemplate.Renderer {
// 	r := multitemplate.NewRenderer()
// 	r.AddFromFiles("index", "web/templates/base.tmpl", "web/templates/index.tmpl")
// 	r.AddFromFiles("login", "web/templates/base.tmpl", "web/templates/login.html")
// 	return r
// }

func main() {
	// fmt.Println("Hello World")
	router := gin.Default()
	router.HTMLRender = loadTemplates("/Users/joshuawelsh/webdev/joshjwelsh.github.io/httpd/templates")
	// router.LoadHTMLGlob("web/templates/*")

	router.GET("/", handler.GetIndex())
	router.GET("/login", handler.GetLogin())
	router.GET("/code")
	router.Run(":3030")
}
