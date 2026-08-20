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
	"github.com/sensorsdata/sa-sdk-go/utils/xrand"
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

// TrackIdGenerator 定义了生成事件 TrackID 的函数类型。
// etype: 事件类型；properties: 事件属性；返回值: TrackID
type TrackIdGenerator func(etype string, properties map[string]interface{}) int32

// defaultTrackIdGenerator 默认的 TrackID 生成器
var defaultTrackIdGenerator = func(etype string, properties map[string]interface{}) int32 {
	return xrand.Int32()
}

type SensorsAnalytics struct {
	C                consumers.Consumer
	ProjectName      string
	TimeFree         bool
	trackIdGenerator TrackIdGenerator // 私有字段，仅通过 Options 设置
}

// Option 定义 SDK 配置选项的函数类型
type Option func(*SensorsAnalytics)

// WithTrackIdGenerator 设置自定义的 TrackID 生成器。
// 必须在 InitSensorsAnalytics 时传入，初始化后不可修改（保证并发安全）。
func WithTrackIdGenerator(gen TrackIdGenerator) Option {
	return func(sa *SensorsAnalytics) {
		if gen != nil {
			sa.trackIdGenerator = gen
		}
	}
}

func InitSensorsAnalytics(c consumers.Consumer, projectName string, timeFree bool, opts ...Option) SensorsAnalytics {
	sa := SensorsAnalytics{
		C:                c,
		ProjectName:      projectName,
		TimeFree:         timeFree,
		trackIdGenerator: defaultTrackIdGenerator,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(&sa)
		}
	}
	return sa
}

func (sa *SensorsAnalytics) Flush() {
	sa.C.Flush()
}

func (sa *SensorsAnalytics) Close() {
	sa.C.Close()
}

// generateTrackID 生成 TrackID。
// 由于 trackIdGenerator 在初始化后只读，此方法天然并发安全。
func (sa *SensorsAnalytics) generateTrackID(etype string, properties map[string]interface{}) int32 {
	if sa.trackIdGenerator != nil {
		return sa.trackIdGenerator(etype, properties)
	}
	return defaultTrackIdGenerator(etype, properties)
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
