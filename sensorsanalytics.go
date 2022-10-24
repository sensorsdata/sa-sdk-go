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

package sensorsanalytics

import (
	"errors"
	"github.com/sensorsdata/sa-sdk-go/consumers"
	"github.com/sensorsdata/sa-sdk-go/utils"
)

const (
	TRACK             = "track"
	TRACK_SIGNUP      = "track_signup"
	PROFILE_SET       = "profile_set"
	PROFILE_SET_ONCE  = "profile_set_once"
	PROFILE_INCREMENT = "profile_increment"
	PROFILE_APPEND    = "profile_append"
	PROFILE_UNSET     = "profile_unset"
	PROFILE_DELETE    = "profile_delete"
	ITEM_SET          = "item_set"
	ITEM_DELETE       = "item_delete"
	MAX_ID_LEN        = 255
)

// 静态公共属性
var superProperties map[string]interface{}

type SensorsAnalytics struct {
	C           consumers.Consumer
	ProjectName string
	TimeFree    bool
}

func InitSensorsAnalytics(c consumers.Consumer, projectName string, timeFree bool) SensorsAnalytics {
	return SensorsAnalytics{C: c, ProjectName: projectName, TimeFree: timeFree}
}

func (sa *SensorsAnalytics) Flush() {
	sa.C.Flush()
}

func (sa *SensorsAnalytics) Close() {
	sa.C.Close()
}

func (sa *SensorsAnalytics) Track(distinctId, event string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	// merge properties
	if properties == nil {
		nproperties = make(map[string]interface{})
	} else {
		nproperties = utils.DeepCopy(properties)
	}

	// merge super properties
	if superProperties != nil {
		utils.MergeSuperProperty(superProperties, nproperties)
	}
	return TrackEvent(sa, TRACK, event, distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) TrackSignup(distinctId, originId string) error {
	// check originId and merge properties
	if originId == "" {
		return errors.New("property [original_id] must not be empty")
	}
	if len(originId) > MAX_ID_LEN {
		return errors.New("the max length of property [original_id] is 255")
	}

	properties := make(map[string]interface{})
	// merge super properties
	if superProperties != nil {
		utils.MergeSuperProperty(superProperties, properties)
	}
	return TrackEvent(sa, TRACK_SIGNUP, "$SignUp", distinctId, originId, properties, false)
}

func (sa *SensorsAnalytics) ProfileSet(distinctId string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEvent(sa, PROFILE_SET, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ProfileSetOnce(distinctId string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEvent(sa, PROFILE_SET_ONCE, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ProfileIncrement(distinctId string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEvent(sa, PROFILE_INCREMENT, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ProfileAppend(distinctId string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEvent(sa, PROFILE_APPEND, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ProfileUnset(distinctId string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEvent(sa, PROFILE_UNSET, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ProfileDelete(distinctId string, isLoginId bool) error {
	nproperties := make(map[string]interface{})
	return TrackEvent(sa, PROFILE_DELETE, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ItemSet(itemType string, itemId string, properties map[string]interface{}) error {
	return ItemTrack(sa, ITEM_SET, itemType, itemId, properties)
}

func (sa *SensorsAnalytics) ItemDelete(itemType string, itemId string) error {
	return ItemTrack(sa, ITEM_DELETE, itemType, itemId, nil)
}

func (sa *SensorsAnalytics) ItemDelete3(itemType string, itemId string, properties map[string]interface{}) error {
	return ItemTrack(sa, ITEM_DELETE, itemType, itemId, properties)
}

// RegisterSuperProperties 注册公共属性
func (sa *SensorsAnalytics) RegisterSuperProperties(superProperty map[string]interface{}) {
	if superProperties == nil {
		superProperties = make(map[string]interface{})
	}
	utils.MergeSuperProperty(superProperty, superProperties)
}

// ClearSuperProperties 清除公共属性
func (sa *SensorsAnalytics) ClearSuperProperties() {
	superProperties = make(map[string]interface{})
}

// UnregisterSuperProperty 清除指定 key 的公共属性
func (sa *SensorsAnalytics) UnregisterSuperProperty(key string) {
	delete(superProperties, key)
}

func InitDefaultConsumer(url string, timeout int) (*consumers.DefaultConsumer, error) {
	return consumers.InitDefaultConsumer(url, timeout)
}

func InitBatchConsumer(url string, max, timeout int) (*consumers.BatchConsumer, error) {
	return consumers.InitBatchConsumer(url, max, timeout)
}

func InitLoggingConsumer(filename string, hourRotate bool) (*consumers.LoggingConsumer, error) {
	return consumers.InitLoggingConsumer(filename, hourRotate)
}

func InitConcurrentLoggingConsumer(filename string, hourRotate bool) (*consumers.ConcurrentLoggingConsumer, error) {
	return consumers.InitConcurrentLoggingConsumer(filename, hourRotate)
}

func InitDebugConsumer(url string, writeData bool, timeout int) (*consumers.DebugConsumer, error) {
	return consumers.InitDebugConsumer(url, writeData, timeout)
}
