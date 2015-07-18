package osm

import (
	"github.com/gardster/ORME/geom"
	"testing"
)

func TestGPXUrlConstructor(t *testing.T) {
	a := geom.New(0, 51.5)
	b := geom.New(0.25, 51.75)
	input := geom.CalculateBBox([]geom.Coordinate{a, b})
	result := GetGPXFetcherUrl(input)
	if result != "http://api.openstreetmap.org/api/0.6/trackpoints?bbox=0,51.5,0.25,51.75" {
		t.Error("Bad formed url: ", result)
	}
}
