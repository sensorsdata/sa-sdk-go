package sensorsanalytics

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/sensorsdata/sa-sdk-go/consumers"
	"github.com/sensorsdata/sa-sdk-go/structs"
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

	SDK_VERSION = "2.0.1"
	LIB_NAME    = "Golang"

	MAX_ID_LEN = 255
)

type SensorsAnalytics struct {
	C           consumers.Consumer
	ProjectName string
	TimeFree    bool
}

func InitSensorsAnalytics(c consumers.Consumer, projectName string, timeFree bool) SensorsAnalytics {
	return SensorsAnalytics{C: c, ProjectName: projectName, TimeFree: timeFree}
}

func (sa *SensorsAnalytics) track(etype, event, distinctId, originId string, properties map[string]interface{}, isLoginId bool) error {
	eventTime := utils.NowMs()
	if et := extractUserTime(properties); et > 0 {
		eventTime = et
	}

	data := structs.EventData{
		Type:          etype,
		Time:          eventTime,
		DistinctId:    distinctId,
		Properties:    properties,
		LibProperties: getLibProperties(),
	}

	if sa.ProjectName != "" {
		data.Project = sa.ProjectName
	}

	if etype == TRACK || etype == TRACK_SIGNUP {
		data.Event = event
	}

	if etype == TRACK_SIGNUP {
		data.OriginId = originId
	}

	if sa.TimeFree {
		data.TimeFree = true
	}

	if isLoginId {
		properties["$is_login_id"] = true
	}

	err := data.NormalizeData()
	if err != nil {
		return err
	}

	return sa.C.Send(data)
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
		properties = make(map[string]interface{})
	} else {
		nproperties = utils.DeepCopy(properties)
	}

	nproperties["$lib"] = LIB_NAME
	nproperties["$lib_version"] = SDK_VERSION

	return sa.track(TRACK, event, distinctId, "", nproperties, isLoginId)
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

	properties["$lib"] = LIB_NAME
	properties["$lib_version"] = SDK_VERSION

	return sa.track(TRACK_SIGNUP, "$SignUp", distinctId, originId, properties, false)
}

func (sa *SensorsAnalytics) ProfileSet(distinctId string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}

	return sa.track(PROFILE_SET, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ProfileSetOnce(distinctId string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}

	return sa.track(PROFILE_SET_ONCE, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ProfileIncrement(distinctId string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}

	return sa.track(PROFILE_INCREMENT, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ProfileAppend(distinctId string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}

	return sa.track(PROFILE_APPEND, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ProfileUnset(distinctId string, properties map[string]interface{}, isLoginId bool) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}

	return sa.track(PROFILE_UNSET, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ProfileDelete(distinctId string, isLoginId bool) error {
	nproperties := make(map[string]interface{})

	return sa.track(PROFILE_DELETE, "", distinctId, "", nproperties, isLoginId)
}

func (sa *SensorsAnalytics) ItemSet(itemType string, itemId string, properties map[string]interface{}) error {
	libProperties := getLibProperties()
	time := utils.NowMs()
	if properties == nil {
		properties = map[string]interface{}{}
	}

	itemData := structs.Item{
		Type:          ITEM_SET,
		ItemId:        itemId,
		Time:          time,
		ItemType:      itemType,
		Properties:    properties,
		LibProperties: libProperties,
	}

	err := itemData.NormalizeItem()
	if err != nil {
		return err
	}

	return sa.C.ItemSend(itemData)
}

func (sa *SensorsAnalytics) ItemDelete(itemType string, itemId string) error {
	libProperties := getLibProperties()
	time := utils.NowMs()

	itemData := structs.Item{
		Type:          ITEM_DELETE,
		ItemId:        itemId,
		Time:          time,
		ItemType:      itemType,
		Properties:    map[string]interface{}{},
		LibProperties: libProperties,
	}

	err := itemData.NormalizeItem()
	if err != nil {
		return err
	}

	return sa.C.ItemSend(itemData)
}

func getLibProperties() structs.LibProperties {
	lp := structs.LibProperties{}
	lp.Lib = LIB_NAME
	lp.LibVersion = SDK_VERSION
	lp.LibMethod = "code"
	if pc, file, line, ok := runtime.Caller(3); ok { //3 means sdk's caller
		f := runtime.FuncForPC(pc)
		lp.LibDetail = fmt.Sprintf("##%s##%s##%d", f.Name(), file, line)
	}

	return lp
}

func extractUserTime(p map[string]interface{}) int64 {
	if t, ok := p["$time"]; ok {
		v, ok := t.(int64)
		if !ok {
			fmt.Fprintln(os.Stderr, "It's not ok for type string")
			return 0
		}
		delete(p, "$time")

		return v
	}

	return 0
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
