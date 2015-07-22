package osm

import (
	"encoding/xml"
	"github.com/gardster/ORME/geom"
	"time"
)

type gpx struct {
	XMLName xml.Name `xml:"gpx"`
	Tracks  []Trk    `xml:"trk"`
}

type Trk struct {
	//XMLName xml.Name `xml:"trk"`
	Segs []TrkSeg `xml:"trkseg"`
}

type TrkSeg struct {
	Points []Trkpt `xml:"trkpt"`
}

type Trkpt struct {
	//XMLName xml.Name `xml:"trkpt"`
	Lat  float32 `xml:"lat,attr"`
	Lon  float32 `xml:"lon,attr"`
	Time string  `xml:"time"`
}

func ParseXML(str []byte) []geom.Track {
	var t gpx
	var result []geom.Track
	err := xml.Unmarshal(str, &t)
	if err != nil {
		panic(err)
	}
	for _, track := range t.Tracks {
		for _, segment := range track.Segs {
			var tr geom.Track
			for _, point := range segment.Points {
				var pt geom.TrackPoint
				pt.Point = geom.Coordinate{point.Lat, point.Lon}
				pt.Time = ParseTime(point.Time)
				tr.Points = append(tr.Points, pt)
			}
			result = append(result, tr)
		}
	}
	return result
}

func ParseTime(timeStr string) int {
	layout := time.RFC3339
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		panic(err)
	}
	return int(t.Unix())
}
