package iterate

// Response represents a response from the API which
// optionally contrains results and error.
type Response struct {
	Results interface{} `json:"results,omitempty"`
	Error   string      `json:"error,omitempty"`
}
