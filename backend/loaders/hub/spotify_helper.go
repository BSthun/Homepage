package hub

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"backend/loaders/mysql"
	"backend/types/enum"
	"backend/types/model"
	"backend/utils/text"
	"backend/utils/value"
)

func spotifyRecord(entity enum.SpotifyEntity, val string) {
	var count uint64
	date := text.GetDateString()
	if result := mysql.DB.Model(new(model.DiarySpotifyAnalytics)).Select("COUNT(*)").Where("date = ? AND entity = ? AND value = ?", date, entity, val).First(&count); result.Error != nil {
		logrus.Warnf("Unable to get spotify analytics (%s): %v", entity, result.Error)
		return
	}
	if count == 0 {
		if result := mysql.DB.Create(&model.DiarySpotifyAnalytics{
			Date:   &date,
			Entity: &entity,
			Value:  &val,
			Count:  value.Ptr[uint64](1),
		}); result.Error != nil {
			logrus.Warnf("Unable to create spotify analytics (%s): %v", entity, result.Error)
			return
		}
	} else {
		if result := mysql.DB.Model(new(model.DiarySpotifyAnalytics)).Where("date = ? AND entity = ? AND value = ?", date, entity, val).Update("count", gorm.Expr("count + ?", 1)); result.Error != nil {
			logrus.Warn("Unable to update spotify analytics (%s): %v", entity, result.Error)
			return
		}
	}
}
