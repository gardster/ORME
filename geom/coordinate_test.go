package geom

import (
	"reflect"
	"testing"
)

func TestCoordinateCreation(t *testing.T) {
	a := Coordinate{35.67, 38.25}
	if a.Lat != 35.67 {
		t.Error("Excepted 35.67, found", a.Lat)
	}
	if a.Lon != 38.25 {
		t.Error("Excepted 38.25, found", a.Lon)
	}
}

func TestCoordinateConstructor(t *testing.T) {
	a := Coordinate{23.34, -123.45}
	b := New(23.34, -123.45)
	if a.Lat != b.Lat || a.Lon != b.Lon {
		t.Error("Constructors are not equal", a, b)
	}
}

func TestCalculateBBox(t *testing.T) {
	bbox := CalculateBBox([]Coordinate{{1, 2}, {3, 4}, {5, 6}})
	if !reflect.DeepEqual(bbox.Min, Coordinate{1, 2}) {
		t.Error("Excepted {1,2}, found", bbox.Min)
	}
	if !reflect.DeepEqual(bbox.Max, Coordinate{5, 6}) {
		t.Error("Excepted {5,6}, found", bbox.Max)
	}
}
