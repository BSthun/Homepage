package diaryEndpoints

import (
	"fmt"
	"share/types/model_extended"
	"share/utils/text"
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/loaders/mysql"
	"backend/types/common"
	"backend/types/response"
	"share/types/payload"
)

func ActivityGraphHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("diary/graph")

	// * Parse query
	query := new(payload.DiaryOverviewQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(false, "Unable to parse query", err)
	}

	// * Validate query
	if err := text.Validator.Struct(query); err != nil {
		return response.Error(false, "Unable to validate query", err)
	}

	// * Set session detail
	s.SetDetail("year", query.Year)

	// * Declare variables
	year := *query.Year

	// * Query diary records
	var diaries []*extendedModel.StrapiDiary
	if result := mysql.StrapiDB.Table("diaries").Select("date, graph, summary").Where("date LIKE ?", fmt.Sprintf("%d-%%", year)).Scan(&diaries); result.Error != nil {
		return response.Error(false, "Unable to query diary records", result.Error)
	}

	// * Construct day mapper
	dayMapper := make(map[string]*extendedModel.StrapiDiary)
	for _, diary := range diaries {
		dayMapper[*diary.Date] = diary
	}

	// * Map year to array
	dayGraph := make([]uint32, 0)
	for iter := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC); iter.Year() == year; iter = iter.AddDate(0, 0, 1) {
		if day, ok := dayMapper[iter.Format("2006-01-02")]; ok {
			lightenGraph := *day.Graph + 0x181818
			if lightenGraph > 0xFFFFFF {
				lightenGraph = 0xFFFFFF
			}
			dayGraph = append(dayGraph, lightenGraph)
		} else if iter.Before(time.Now()) {
			dayGraph = append(dayGraph, 0x121212)
		} else {
			dayGraph = append(dayGraph, 0x181818)
		}
	}

	// * Commit session
	_ = s.Commit(nil)

	return c.JSON(response.Success(map[string]any{
		"days_graph": dayGraph,
	}))
}
