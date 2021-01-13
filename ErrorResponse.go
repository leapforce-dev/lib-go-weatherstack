package weatherstack

// ErrorResponse stores general Ridder API error response
//
type ErrorResponse struct {
	Success string `json:"success"`
	Error   struct {
		Code int    `json:"code"`
		Type string `json:"type"`
		Info string `json:"info"`
	} `json:"error"`
}
