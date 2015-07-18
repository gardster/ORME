package geom

type Coordinate struct {
	Lat float32
	Lon float32
}

type BBox struct {
	Min Coordinate
	Max Coordinate
}

func New(lat float32, lon float32) Coordinate {
	return Coordinate{lat, lon}
}

func CalculateBBox(coords []Coordinate) BBox {
	minCoord := coords[0]
	maxCoord := coords[0]
	for _, v := range coords {
		if v.Lat < minCoord.Lat {
			minCoord.Lat = v.Lat
		}
		if v.Lon < minCoord.Lon {
			minCoord.Lon = v.Lon
		}
		if v.Lat > maxCoord.Lat {
			maxCoord.Lat = v.Lat
		}
		if v.Lon > maxCoord.Lon {
			maxCoord.Lon = v.Lon
		}
	}
	return BBox{minCoord, maxCoord}
}
