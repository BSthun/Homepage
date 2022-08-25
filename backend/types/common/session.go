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
	Id uint64 `json:"id"`
}

func (r *Session) Log(tag enum.Tag, detail map[string]any) *response.ErrorInstance {
	bytes, err := json.Marshal(detail)
	if err != nil {
		return &response.ErrorInstance{
			Message: "Failed to marshal session log detail",
		}
	}

	log := &model.TrackLog{
		TrackSessionId: &r.Id,
		Tag:            &tag,
		Detail:         value.Ptr(string(bytes)),
		CreatedAt:      nil,
		UpdatedAt:      nil,
	}
	if result := mysql.DB.Create(log); result.Error != nil {
		return response.Error(true, "Failed to log session", result.Error)
	}

	return nil
}
