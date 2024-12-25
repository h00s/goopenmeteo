package goopenmeteo

type ErrorResponse struct {
	Error  bool   `json:"error"`
	Reason string `json:"reason"`
}
