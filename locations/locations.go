package locations

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Locations struct {
	Locations []Location
}

type Location struct {
	Id          int
	Title       string
	Content     string
	Opened      bool
	Mask        string
	Towel       string
	Fountain    string
	Locker_room string
	Schedules   []Schedule
}

type Schedule struct {
	Weekdays string
	Hour     string
}

type SmallerLocation struct {
	Id         int
	Title      string
	Street     string
	Region     string
	City_name  string
	State_name string
	Uf         string
}

func GetLocations() []Location {
	locations_json, err := os.ReadFile("locations.json")

	if err != nil {
		panic(err)
	}

	var locations Locations

	err = json.Unmarshal(locations_json, &locations)
	if err != nil {
		panic(err)
	}

	return locations.Locations
}

func FilterLocations(locations []Location, lower_bound_hr int, lower_bound_min int, upper_bound_hr int, upper_bound_min int, show_closed bool) []Location {
	var filtered []Location

	for _, location := range locations {
		if !show_closed && !location.Opened {
			continue
		}

		for _, schedule := range location.Schedules {
			hours := strings.Split(schedule.Hour, " ")

			if len(hours) < 2 {
				continue
			}

			var first_hour, first_min int
			_, err := fmt.Sscanf(hours[0], "%dh%d", &first_hour, &first_min)
			if err != nil {
				_, err := fmt.Sscanf(hours[0], "%dh", &first_hour)

				first_min = 0
				if err != nil {
					log.Println(err)
				}
			}

			var second_hour, second_min int
			_, err2 := fmt.Sscanf(hours[2], "%dh%d", &second_hour, &second_min)
			if err2 != nil {
				_, err := fmt.Sscanf(hours[2], "%dh", &second_hour)

				first_min = 0
				if err != nil {
					log.Println(err)
				}
			}

			if is_earlier(first_hour, first_min, lower_bound_hr, lower_bound_min) && is_later(second_hour, second_min, lower_bound_hr, lower_bound_min) {
				filtered = append(filtered, location)
				break
			} else if is_earlier(first_hour, first_min, upper_bound_hr, upper_bound_min) && is_later(second_hour, second_min, upper_bound_hr, upper_bound_min) {
				filtered = append(filtered, location)
				break
			}
		}
	}

	return filtered
}

func is_earlier(hr1 int, min1 int, hr2 int, min2 int) bool {
	if hr1 < hr2 {
		return true
	} else if hr1 == hr2 && min1 < min2 {
		return true
	}

	return false
}

func is_later(hr1 int, min1 int, hr2 int, min2 int) bool {
	if hr1 > hr2 {
		return true
	} else if hr1 == hr2 && min1 > min2 {
		return true
	}

	return false
}
