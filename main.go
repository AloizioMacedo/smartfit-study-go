package main

import (
	"fmt"
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

func get_source(requirement string, obj string) string {
	var req string
	if requirement == "allowed" {
		req = "required"
	} else if requirement == "not_allowed" || requirement == "closed" {
		req = "forbidden"
	} else {
		req = requirement
	}

	return fmt.Sprintf("images/%s-%s.png", req, obj)
}

func get_alt(requirement string, obj string) string {
	return fmt.Sprintf("%s-%s", requirement, obj)
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

		var open_class string
		var opened_status string
		if location.Opened {
			open_class = "open-facility"
			opened_status = "Aberto"
		} else {
			open_class = "closed-facility"
			opened_status = "Fechado"
		}

		results = append(results, Result{
			OpenClass:    open_class,
			OpenedStatus: opened_status,
			Title:        location.Title,
			Address:      template.HTML(location.Content),
			Prohibs: []Prohib{
				{
					ProhibSource: get_source(location.Mask, "mask"),
					Alt:          get_alt(location.Mask, "mask"),
				}, {
					ProhibSource: get_source(location.Towel, "towel"),
					Alt:          get_alt(location.Towel, "towel"),
				},
				{
					ProhibSource: get_source(location.Fountain, "fountain"),
					Alt:          get_alt(location.Fountain, "fountain"),
				},
				{
					ProhibSource: get_source(location.Locker_room, "lockerroom"),
					Alt:          get_alt(location.Locker_room, "lockerroom"),
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
