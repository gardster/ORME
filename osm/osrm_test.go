package osm

import (
	"github.com/gardster/ORME/geom"
	"testing"
)

func TestUrlGenerator(t *testing.T) {
	master := "https://router.project-osrm.org/match?loc=52.542648,13.393252&t=1424684612&loc=52.543079,13.394780&t=1424684616&loc=52.542107,13.397389&t=1424684620"
	points := []geom.TrackPoint{geom.TrackPoint{geom.Coordinate{52.542648, 13.393252}, 1424684612}, geom.TrackPoint{geom.Coordinate{52.543079, 13.394780}, 1424684616}, geom.TrackPoint{geom.Coordinate{52.542107, 13.397389}, 1424684620}}
	input := geom.Track{points}
	result := MakeMatchingUrl(input)
	if master != result {
		t.Error("URL does not match given pattern", result)
	}
}
