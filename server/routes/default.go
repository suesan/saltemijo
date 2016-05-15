package routes

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/flosch/pongo2.v3"
)

func Root(c *gin.Context) {
	tpl, err := pongo2.FromFile("index.html")
	if err != nil {
		c.String(500, "Internal Server Error.")
	}
	err = tpl.ExecuteWriter(pongo2.Context{}, c.Writer)
	if err != nil {
		c.String(500, "Internal Server Error.")
	}
}
