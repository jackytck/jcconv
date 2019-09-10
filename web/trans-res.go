package web

// TransRes represents the JSON response.
type TransRes struct {
	Input  string `json:"input"`
	Output string `json:"output"`
	Error  string `json:"error"`
}
