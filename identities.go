// Package sensorsanalytics /*
package sensorsanalytics

import (
	"errors"
	"github.com/sensorsdata/sa-sdk-go/utils"
)

const (
	BIND         = "track_id_bind"
	UNBIND       = "track_id_unbind"
	BIND_EVENT   = "$BindID"
	UNBIND_EVENT = "$UnbindID"
	LOGIN_ID     = "$identity_login_id"
	MOBILE       = "$identity_mobile"
	EMAIL        = "$identity_email"
)

type Identity struct {
	Identities map[string]string
}

func (sa *SensorsAnalytics) Bind(identity Identity) error {
	if identity.Identities == nil || len(identity.Identities) < 2 {
		return errors.New("identity is invalid")
	}
	return TrackEventID3(sa, identity, BIND, BIND_EVENT, nil)
}

func (sa *SensorsAnalytics) UnBind(identity Identity) error {
	if identity.Identities == nil {
		return errors.New("identity is nil")
	}
	return TrackEventID3(sa, identity, UNBIND, UNBIND_EVENT, nil)
}

func (sa *SensorsAnalytics) TrackById(identity Identity, event string, properties map[string]interface{}) error {
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

	return TrackEventID3(sa, identity, TRACK, event, nproperties)
}

func (sa *SensorsAnalytics) ProfileSetById(identity Identity, properties map[string]interface{}) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEventID3(sa, identity, PROFILE_SET, "", nproperties)
}

func (sa *SensorsAnalytics) ProfileSetOnceById(identity Identity, properties map[string]interface{}) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEventID3(sa, identity, PROFILE_SET_ONCE, "", nproperties)
}

func (sa *SensorsAnalytics) ProfileIncrementById(identity Identity, properties map[string]interface{}) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEventID3(sa, identity, PROFILE_INCREMENT, "", nproperties)
}

func (sa *SensorsAnalytics) ProfileAppendById(identity Identity, properties map[string]interface{}) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEventID3(sa, identity, PROFILE_APPEND, "", nproperties)
}

func (sa *SensorsAnalytics) ProfileUnsetById(identity Identity, properties map[string]interface{}) error {
	var nproperties map[string]interface{}

	if properties == nil {
		return errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEventID3(sa, identity, PROFILE_UNSET, "", nproperties)
}

func (sa *SensorsAnalytics) ProfileDeleteById(identity Identity) error {
	nproperties := make(map[string]interface{})
	return TrackEventID3(sa, identity, PROFILE_DELETE, "", nproperties)
}
