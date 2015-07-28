package osm

import (
	"fmt"
	"github.com/gardster/ORME/geom"
)

const (
	SERVER = "https://router.project-osrm.org"
)

func MakeMatchingUrl(track geom.Track) string {
	resultStr := ""
	for _, v := range track.Points {
		resultStr = fmt.Sprintf("%s&loc=%.6f,%.6f&t=%d", resultStr, v.Point.Lat, v.Point.Lon, v.Time)
	}
	return fmt.Sprintf("%s/match?%s", SERVER, resultStr[1:])
}
