package geom

import ()

type TrackPoint struct {
	Point Coordinate
	Time  int
}

type Track struct {
	Points []TrackPoint
}
