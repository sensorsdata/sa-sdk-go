/*
 * Created by dengshiwei on 2020/01/06.
 * Copyright 2015－2020 Sensors Data Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"github.com/sensorsdata/sa-sdk-go/structs"
)

var typemap map[string]int = map[string]int{
	"track":             0,
	"track_signup":      0,
	"profile_set":       0,
	"profile_set_once":  0,
	"profile_increment": 0,
	"profile_append":    0,
	"profile_unset":     0,
	"profile_delete":    0,
}

var p map[string]interface{} = map[string]interface{}{
	"$ip":            "1.1.1.1",
	"$is_login_id":   true,
	"$lib":           "Golang",
	"$lib_version":   "1.7.5",
	"ProductCatalog": "Laptop Computer",
	"ProductId":      "123456",
}

var demoLibData map[string]interface{} = map[string]interface{}{
	"$lib":         "Golang",
	"$lib_version": "1.7.5",
	"$lib_method":  "code",
	"$lib_detail":  "##main.main.go##track.go##28",
}
var demoEventData map[string]interface{} = map[string]interface{}{
	"type":        "track",
	"time":        DemoTime,
	"distinct_id": DemoDistinctId,
	"properties":  p,
	"lib":         demoLibData,
	"project":     "default",
}

var DemoDistinctId string = "ABCDEF123456"
var DemoEventString string = "ViewProduct"
var DemoTime int64 = 1523329458000
var DemoProperties map[string]interface{} = map[string]interface{}{
	"$ip":            "1.1.1.1",
	"ProductId":      "123456",
	"ProductCatalog": "Laptop Computer",
	"IsAddedToFav":   true,
	"$is_login_id":   true,
	"$lib":           "Golang",
	"$lib_version":   "1.7.5",
}

var DemoLib structs.LibProperties = structs.LibProperties{
	Lib:        "Golang",
	LibVersion: "1.7.5",
	LibMethod:  "code",
	AppVersion: "",
	LibDetail:  "##main.main.go##track.go##28",
}

var DemoEvent structs.EventData = structs.EventData{
	Type:          "track",
	Time:          DemoTime,
	DistinctId:    DemoDistinctId,
	Properties:    DemoProperties,
	LibProperties: DemoLib,
	Project:       "default",
}

func demoCompare(ed structs.EventData) string {
	if ed.DistinctId != DemoDistinctId {
		return "distinctid error"
	}
	if ed.Time != DemoEvent.Time {
		return "time error"
	}
	if ed.Project != DemoEvent.Project {
		return "project error"
	}
	if _, ok := typemap[ed.Type]; !ok {
		return "type error"
	}
	lib := ed.LibProperties
	if lib.Lib != DemoLib.Lib {
		return "lib.lib error"
	}
	if lib.LibVersion != DemoLib.LibVersion {
		return "lib.lib_version error"
	}
	if lib.LibMethod != DemoLib.LibMethod {
		return "lib.lib_method error"
	}
	if lib.LibDetail == "" {
		return "lib.lib_detail error"
	}
	properties := ed.Properties
	if properties["$ip"] != DemoEvent.Properties["$ip"] {
		return "properties.$ip error"
	}
	if properties["$is_login_id"] != DemoEvent.Properties["$is_login_id"] {
		return "properties.$is_login_id error"
	}
	if properties["$lib"] != DemoEvent.Properties["$lib"] {
		return "properties.$lib error"
	}
	if properties["$lib_version"] != DemoEvent.Properties["$lib_version"] {
		return "properties.$lib_version error"
	}
	if properties["ProductId"] != DemoEvent.Properties["ProductId"] {
		return "properties.ProductId error"
	}

	return ""
}

func demoListCompare(edl EDL) string {
	for _, ed := range edl {
		estr := demoCompare(ed)
		if estr != "" {
			return estr
		}
	}
	return ""
}
