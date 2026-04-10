package transport

type CreateRequest struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
