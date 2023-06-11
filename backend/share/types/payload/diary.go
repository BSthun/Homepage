package payload

type DiaryOverviewQuery struct {
	Year *int `json:"year_no" validate:"min=2022,max=2100"`
}
