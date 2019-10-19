package web

// TransRes represents the JSON response.
type TransRes struct {
	Input      string `json:"input"`
	FromLocale string `json:"fromLocale"`
	ToLocale   string `json:"toLocale"`
	Chain      string `json:"chain"`
	Output     string `json:"output"`
	Error      string `json:"error"`
	Elapsed    string `json:"elapsed"`
}
