package main

import "github.com/gin-gonic/gin"

type Result struct {
	OpenClass    string
	OpenedStatus string
	Title        string
	Address      string
	Prohibs      []Prohib
	Schedules    []Schedule
}

type Prohib struct {
	ProhibSource string
	Alt          string
}

type Schedule struct {
	Weekdays string
	Hour     string
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func main() {
	r := gin.Default()
	r.StaticFile("style.css", "./templates/style.css")
	r.Static("images", "./templates/images")

	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", index)

	r.Run()
}
