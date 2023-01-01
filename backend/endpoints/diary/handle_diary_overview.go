package diaryEndpoints

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/loaders/mysql"
	"backend/types/common"
	"backend/types/model_extended"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
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
			lightenGraph := *day.Graph + 1579032
			if lightenGraph > 16777215 {
				lightenGraph = 16777215 /* #FFFFFF */
			}
			dayGraph = append(dayGraph, lightenGraph)
		} else if iter.Before(time.Now()) {
			dayGraph = append(dayGraph, 1184274 /* #121212 */)
		} else {
			dayGraph = append(dayGraph, 1579032 /* #181818 */)
		}
	}

	// * Commit session
	_ = s.Commit(nil)

	return c.JSON(response.Info(map[string]any{
		"days_graph": dayGraph,
	}))
}
