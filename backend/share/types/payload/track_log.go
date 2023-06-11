package payload

type LogBody struct {
	Event          string `json:"event"`
	BeginningState any    `json:"beginning_state"`
	EndingState    any    `json:"ending_state"`
}

type SigBody struct {
	Title string `json:"title"`
}
