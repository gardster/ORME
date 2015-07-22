package osm

import (
	"encoding/xml"
	"fmt"
	"github.com/gardster/ORME/geom"
	"time"
)

type trk struct {
	XMLName xml.Name `xml:"trkpt"`
	Lat     float64  `xml:"lat,attr"`
	Lon     float64  `xml:"lon,attr"`
	Time    string   `xml:"time"`
}

func ParseXML(str []byte) geom.Coordinate {
	var t trk
	err := xml.Unmarshal(str, &t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Time)
	return geom.Coordinate{float32(t.Lat), float32(t.Lon)}
}

func ParseTime(timeStr string) int64 {
	layout := time.RFC3339
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		panic(err)
	}
	return int64(t.Unix())
}
