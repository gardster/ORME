package main

type Coordinate struct {
	lat float32
	lon float32
}

type BBox struct {
	min Coordinate
	max Coordinate
}

func CalculateBBox(coords []Coordinate) BBox {
	minCoord := coords[0]
	maxCoord := coords[0]
	for _, v := range coords {
		if v.lat < minCoord.lat {
			minCoord.lat = v.lat
		}
		if v.lon < minCoord.lon {
			minCoord.lon = v.lon
		}
		if v.lat > maxCoord.lat {
			maxCoord.lat = v.lat
		}
		if v.lon > maxCoord.lon {
			maxCoord.lon = v.lon
		}
	}
	return BBox{minCoord, maxCoord}
}
