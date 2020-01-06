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
	"regexp"
	"time"
)

type Item struct {
	Type          string                 `json:"type"`
	Time          int64                  `json:"time"`
	ItemType      string                 `json:"item_type"`
	ItemId        string                 `json:"item_id"`
	Properties    map[string]interface{} `json:"properties"`
	LibProperties LibProperties          `json:"lib"`
}

func init() {
	patternBad, _ = regexp.Compile(NAME_PATTERN_BAD)
	patternOk, _ = regexp.Compile(NAME_PATTERN_OK)
}

func (item *Item) NormalizeItem() error {
	// check type
	if item.Type == "" || len(item.Type) == 0 {
		return errors.New("ItemType can't be empty")
	}

	if !checkPattern([]byte(item.Type)) {
		return errors.New("The key '" + item.Type + "' is invalid.")
	}

	if len(item.Type) > 255 {
		return errors.New("The key '" + item.Type + "' is too long, max length is 255.")
	}

	//check item_id
	if item.ItemId == "" || len(item.ItemId) == 0 {
		return errors.New("ItemId must not be empty.")
	}

	if len(item.ItemId) > 255 {
		return errors.New("The ItemId '" + item.ItemId + "' is too long, max length is 255.")
	}

	//check properties
	if item.Properties != nil {
		for k, v := range item.Properties {
			//check key
			if len(k) > KEY_MAX {
				return errors.New("the max length of property key is 256")
			}
			isMatch := checkPattern([]byte(k))
			if !isMatch {
				return errors.New("property key must be a valid variable name.")
			}

			//check value
			switch v.(type) {
			case int:
			case bool:
			case float64:
			case string:
				if len(v.(string)) > VALUE_MAX {
					return errors.New("the max length of property key is 8192")
				}
			case []string: //value in properties list MUST be string
			case time.Time: //only support time.Time
				item.Properties[k] = v.(time.Time).Format("2006-01-02 15:04:05.999")

			default:
				return errors.New("property value must be a string/int/float64/bool/time.Time/[]string")
			}
		}
	}

	return nil
}
