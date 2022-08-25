package common

import (
	"encoding/json"

	"backend/loaders/mysql"
	"backend/types/enum"
	"backend/types/model"
	"backend/types/response"
	"backend/utils/value"
)

type Session struct {
	Id          uint64         `json:"id"`
	TrackTag    enum.Tag       `json:"track_tag"`
	TrackDetail map[string]any `json:"track_detail"`
}

func (r *Session) SetTag(tag enum.Tag) {
	r.TrackTag = tag
}

func (r *Session) SetDetail(key string, value any) {
	r.TrackDetail[key] = value
}

func (r *Session) Commit(err error) *response.ErrorInstance {
	log := &model.TrackLog{
		TrackSession:   nil,
		TrackSessionId: &r.Id,
		Tag:            &r.TrackTag,
		Detail:         nil,
		Error:          nil,
		CreatedAt:      nil,
		UpdatedAt:      nil,
	}

	if bytes, err := json.Marshal(r.TrackDetail); err != nil {
		return response.Error(true, "Failed to marshal session log detail", err)
	} else {
		log.Detail = value.Ptr(string(bytes))
	}

	if err != nil {
		log.Error = value.Ptr(err.Error())
	}

	if result := mysql.DB.Create(log); result.Error != nil {
		return response.Error(true, "Failed to log session", result.Error)
	}

	return nil
}
