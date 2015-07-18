package osm

import (
	"fmt"
	"github.com/gardster/ORME/geom"
)

func GetGPXFetcherUrl(bbox geom.BBox) string {
	return fmt.Sprintf("http://api.openstreetmap.org/api/0.6/trackpoints?bbox=%v,%v,%v,%v", bbox.Min.Lat, bbox.Min.Lon, bbox.Max.Lat, bbox.Max.Lon)
}
