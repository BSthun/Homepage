package payload

type LogBody struct {
	Event          string `json:"event"`
	BeginningState string `json:"beginning_state"`
	EndingState    string `json:"ending_state"`
}
