package response

type ErrorResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}
