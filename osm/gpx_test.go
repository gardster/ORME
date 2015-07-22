package osm

import (
	"github.com/gardster/ORME/geom"
	"reflect"
	"testing"
)

func TestXMLParsing(t *testing.T) {
	result := ParseXML([]byte(`<trkpt lat="51.6626584" lon="0.2244968"><time>2015-07-19T14:53:53Z</time></trkpt>`))
	if !reflect.DeepEqual(result, geom.Coordinate{51.6626584, 0.2244968}) {
		t.Error("Coordinates are not match", result)
	}
}

func TestTimeParser(t *testing.T) {
	timeStr := "2015-07-19T14:53:53Z"
	result := ParseTime(timeStr)
	if result != 1437317633 {
		t.Error("Parsed time does not match", result)
	}

}
