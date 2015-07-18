package main

type Coordinate struct {
	lat float32
	lon float32
}

type BBox struct {
	mincoord Coordinate
	maxcoord Coordinate
}

func CalculateBBox(coords []Coordinate) BBox {
	mincoord := coords[0]
	maxcoord := coords[0]
	for _, v := range coords {
		if v.lat < mincoord.lat {
			mincoord.lat = v.lat
		}
		if v.lon < mincoord.lon {
			mincoord.lon = v.lon
		}
		if v.lat > maxcoord.lat {
			maxcoord.lat = v.lat
		}
		if v.lon > maxcoord.lon {
			maxcoord.lon = v.lon
		}
	}
	return BBox{mincoord, maxcoord}
}
