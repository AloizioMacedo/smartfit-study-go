package locations

import (
	"encoding/json"
	"os"
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
