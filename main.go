package main

import (
	"html/template"
	"smartfit/locations"

	"github.com/gin-gonic/gin"
)

type Results struct {
	Results []Result
}

type Result struct {
	OpenClass    string
	OpenedStatus string
	Title        string
	Address      template.HTML
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

func parse_locations(locations []locations.Location) Results {
	var results []Result

	for _, location := range locations {
		var schedules []Schedule

		for _, schedule := range location.Schedules {
			schedules = append(schedules, Schedule{
				Weekdays: schedule.Weekdays,
				Hour:     schedule.Hour,
			})
		}

		results = append(results, Result{
			OpenClass:    "Closed",
			OpenedStatus: "Closed",
			Title:        location.Title,
			Address:      template.HTML(location.Content),
			Prohibs: []Prohib{
				{
					ProhibSource: "Closed",
					Alt:          "Closed",
				},
			},
			Schedules: schedules,
		})
	}

	return Results{Results: results}
}

func results(c *gin.Context) {
	// day_period := c.Query("day_period")
	// show_closed := c.DefaultQuery("show_closed", "false")
	locations := locations.GetLocations()

	c.HTML(200, "results.html", parse_locations(locations))
}

func main() {
	r := gin.Default()
	r.StaticFile("style.css", "./templates/style.css")
	r.Static("images", "./templates/images")
	r.Static("fonts", "./templates/fonts")

	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", index)
	r.GET("/results", results)
	r.GET("/clean", clean)

	r.Run()
}
