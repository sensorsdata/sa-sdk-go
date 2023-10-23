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

package structs

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

const (
	KEY_MAX   = 100
	VALUE_MAX = 8192

	NAME_PATTERN_BAD = "^(^distinct_id$|^original_id$|^time$|^properties$|^id$|^first_id$|^second_id$|^users$|^events$|^event$|^user_id$|^date$|^datetime$|^user_group.*|^user_tag.*)$"
	NAME_PATTERN_OK  = "^[a-zA-Z_$][a-zA-Z\\d_$]{0,99}$"
)

var patternBad, patternOk *regexp.Regexp

type EventData struct {
	Type          string                 `json:"type"`
	TrackID       int32                  `json:"_track_id"`
	Time          int64                  `json:"time"`
	DistinctId    string                 `json:"distinct_id,omitempty"`
	Properties    map[string]interface{} `json:"properties"`
	LibProperties LibProperties          `json:"lib"`
	Project       string                 `json:"project"`
	Event         string                 `json:"event"`
	OriginId      string                 `json:"original_id,omitempty"`
	TimeFree      bool                   `json:"time_free,omitempty"`
	Identities    map[string]string      `json:"identities,omitempty"`
}

func init() {
	patternBad, _ = regexp.Compile(NAME_PATTERN_BAD)
	patternOk, _ = regexp.Compile(NAME_PATTERN_OK)
}

func (e *EventData) NormalizeData() error {
	//check distinct id
	if e.DistinctId == "" || len(e.DistinctId) == 0 {
		return errors.New("property [distinct_id] can't be empty")
	}

	if len(e.DistinctId) > 255 {
		return errors.New("the max length of property [distinct_id] is 255")
	}

	//check event
	if e.Event != "" {
		isMatch := checkPattern([]byte(e.Event))
		if !isMatch {
			return errors.New("event name = " + e.Event + " is invalid, event name must be a valid variable name.")
		}
	}

	//check project
	if e.Project != "" {
		isMatch := checkPattern([]byte(e.Project))
		if !isMatch {
			return errors.New("project = " + e.Project + " is invalid, project name must be a valid variable name.")
		}
	}

	//check properties
	if e.Properties != nil {
		for k, v := range e.Properties {
			//check key
			if len(k) > KEY_MAX {
				return errors.New("the max length of property key is 100," + "key = " + k)
			}

			if len(k) == 0 {
				return errors.New("The key is empty or null," + "key = " + k + ", value = " + v.(string))
			}
			isMatch := checkPattern([]byte(k))
			if !isMatch {
				return errors.New("property key must be a valid variable name," + "key = " + k)
			}

			//check value
			switch v.(type) {
			case int:
			case bool:
			case float64:
			case string:
				if len(v.(string)) > VALUE_MAX {
					return errors.New("the max length of property value is 8192," + "value = " + v.(string))
				}
			case []string: //value in properties list MUST be string
			case time.Time: //only support time.Time
				e.Properties[k] = v.(time.Time).Format("2006-01-02 15:04:05.999")

			default:
				return errors.New("property value must be a string/int/float64/bool/time.Time/[]string," + "key = " + k)
			}
		}
	}

	return nil
}

func (e *EventData) CheckIdentities() error {
	if e.Identities == nil || len(e.Identities) < 1 {
		return errors.New("identity is nil")
	}

	for k, v := range e.Identities {
		if len(k) == 0 {
			return errors.New(fmt.Sprintf("The identity invalid, key is empty or null, value = %v", v))
		}

		if !checkPattern([]byte(k)) {
			return errors.New(fmt.Sprintf("The identity invalid, key is preset key, key = %s, value = %v", k, v))
		}

		if len(v) == 0 || len(v) > 255 {
			return errors.New(fmt.Sprintf("The identity invalid, value is empty or length greater than 255, key = %s, value = %v", k, v))
		}
	}
	return nil
}

func checkPattern(name []byte) bool {
	return !patternBad.Match(name) && patternOk.Match(name)
}
