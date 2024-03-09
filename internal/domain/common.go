package domain

type CommonResponse struct {
	Message    string  `json:"message"`
	Error      *string `json:"error"`
	StatusCode int     `json:"status_code"`
}
