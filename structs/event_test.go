package structs

import (
	//"fmt"
	"time"
	"testing"
)

var lib LibProperties = LibProperties {
	Lib        : "Golang",
	LibVersion : "1.1.1",
	LibMethod  : "code",
	AppVersion : "1.0.1",
	LibDetail  : "somedetail",
}

func TestEventBadDistinctId(t *testing.T) {
	//empty distinctid
	ed := EventData {
		Type: "track",
		Time: time.Now().UnixNano() / 1000000,
		DistinctId: "",
		Properties: map[string]interface{}{
			"name": "alice",
			"age": 1,
		},
		LibProperties: lib,
		Project: "default",
		Event: "test",
	}

	if err := ed.NormalizeData(); err != nil {
		t.Log("found bad type", err)
	} else {
		t.Fatal("not found bad type")
	}

	//too large distinctid
	id := make([]byte, 256)
	id[255] = byte('a')
	ed.DistinctId = string(id)

	if err := ed.NormalizeData(); err != nil {
		t.Log("found bad type", err)
	} else {
		t.Fatal("not found bad type")
	}
}

func TestEventBadEvent(t *testing.T) {
	//reserve event name
	ed := EventData {
		Type: "track",
		Time: time.Now().UnixNano() / 1000000,
		DistinctId: "123123",
		Properties: map[string]interface{}{
			"name": "alice",
			"age": 1,
		},
		LibProperties: lib,
		Project: "default",
		Event: "time",
	}

	if err := ed.NormalizeData(); err != nil {
		t.Log("found bad type", err)
	} else {
		t.Fatal("not found bad type")
	}

	//bad event name pattern
	ed.Event = "@$%^abc"
	if err := ed.NormalizeData(); err != nil {
		t.Log("found bad type", err)
	} else {
		t.Fatal("not found bad type")
	}
}

func TestEventBadProject(t *testing.T) {
	ed := EventData {
		Type: "track",
		Time: time.Now().UnixNano() / 1000000,
		DistinctId: "123123",
		Properties: map[string]interface{}{
			"name": "alice",
			"age": 1,
		},
		LibProperties: lib,
		Project: "time",
		Event: "sometime",
	}

	if err := ed.NormalizeData(); err != nil {
		t.Log("found bad type", err)
	} else {
		t.Fatal("not found bad type")
	}

	ed.Project = "@$%^abc"
	if err := ed.NormalizeData(); err != nil {
		t.Log("found bad type", err)
	} else {
		t.Fatal("not found bad type")
	}
}

func TestEventBadPropertiesKey(t *testing.T) {
	//bad pattern
	ed := EventData {
		Type: "track",
		Time: time.Now().UnixNano() / 1000000,
		DistinctId: "123123",
		Properties: map[string]interface{}{
			"name": "alice",
			"age": 1,
			"time": "12345",
		},
		LibProperties: lib,
		Project: "test",
		Event: "sometime",
	}

	if err := ed.NormalizeData(); err != nil {
		t.Log("found bad type", err)
	} else {
		t.Fatal("not found bad type")
	}

	delete(ed.Properties, "time")

	//bad value type
	ed.Properties["abc"] = int64(1)

	if err := ed.NormalizeData(); err != nil {
		t.Log("found bad type", err)
	} else {
		t.Fatal("not found bad type")
	}

	//bad value array type
	badv := []int{1, 2}
	ed.Properties["abc"] = badv
	if err := ed.NormalizeData(); err != nil {
		t.Log("found bad type", err)
	} else {
		t.Fatal("not found bad type")
	}
}

func TestEventOkPropertiesKey(t *testing.T) {
	//good pattern
	ed := EventData {
		Type: "track",
		Time: time.Now().UnixNano() / 1000000,
		DistinctId: "123123",
		Properties: map[string]interface{}{
			"name": "alice",
			"age": 1,
			"size": 0.2,
			"flag": true,
			"sometime": time.Now(),
			"arr": []string{"a", "b"},
		},
		LibProperties: lib,
		Project: "test",
		Event: "sometime",
	}

	if err := ed.NormalizeData(); err != nil {
		t.Fatal("found bad type", err)
	} else {
		t.Log("all type is valid")
	}
}
