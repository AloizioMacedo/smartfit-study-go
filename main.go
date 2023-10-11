package main

import "github.com/gin-gonic/gin"

type Results struct {
	Results []Result
}

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

func clean(c *gin.Context) {
	c.HTML(200, "results.html", Results{[]Result{}})
}

func results(c *gin.Context) {
	// day_period := c.Query("day_period")
	// show_closed := c.DefaultQuery("show_closed", "false")

	c.HTML(200, "results.html", Results{
		Results: []Result{
			{
				OpenClass:    "closed",
				OpenedStatus: "Closed",
				Title:        "Closed",
				Address:      "Closed",
				Prohibs: []Prohib{
					{
						ProhibSource: "Closed",
						Alt:          "Closed",
					},
				},
				Schedules: []Schedule{
					{
						Weekdays: "Closed",
						Hour:     "Closed",
					},
				},
			},
		},
	})
}

func main() {
	r := gin.Default()
	r.StaticFile("style.css", "./templates/style.css")
	r.Static("images", "./templates/images")

	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", index)
	r.GET("/results", results)
	r.GET("/clean", clean)

	r.Run()
}
