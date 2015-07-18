package geom

import (
	"reflect"
	"testing"
)

func TestCoordinateCreation(t *testing.T) {
	a := Coordinate{35.67, 38.25}
	if a.lat != 35.67 {
		t.Error("Excepted 35.67, found", a.lat)
	}
	if a.lon != 38.25 {
		t.Error("Excepted 38.25, found", a.lon)
	}
}

func TestCalculateBBox(t *testing.T) {
	bbox := CalculateBBox([]Coordinate{{1, 2}, {3, 4}, {5, 6}})
	if !reflect.DeepEqual(bbox.min, Coordinate{1, 2}) {
		t.Error("Excepted {1,2}, found", bbox.min)
	}
	if !reflect.DeepEqual(bbox.max, Coordinate{5, 6}) {
		t.Error("Excepted {5,6}, found", bbox.max)
	}
}
