package viewmodels

// PostRequest defines the structure of the request
// payload to create posts entries
type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
