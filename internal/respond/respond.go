package respond

// BadRequest is used to handle API requests with errors
type BadRequest struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
