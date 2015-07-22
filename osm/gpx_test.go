package osm

import (
	"io/ioutil"
	"testing"
)

func TestXMLParsing(t *testing.T) {
	data, _ := ioutil.ReadFile("tracks.gpx")
	result := ParseXML(data)
	if len(result) != 2 {
		t.Error("Can't find 2 tracks in file", len(result))
	}
	if len(result[0].Points) != 560 {
		t.Error("First track contains only", len(result[0].Points), "points")
	}
	if len(result[1].Points) != 4440 {
		t.Error("Second track contains only", len(result[1].Points), "points")
	}
}

func TestTimeParser(t *testing.T) {
	timeStr := "2015-07-19T14:53:53Z"
	result := ParseTime(timeStr)
	if result != 1437317633 {
		t.Error("Parsed time does not match", result)
	}

}
