package requests

type BookRequest struct {
	Title     string `json:"title"`
	Pages     uint64 `json:"pages"`
	WordCount uint64 `json:"word_count"`
	Status    string `json:"status"`
	AuthorID  uint64 `json:"author_id"`
}

type BookStatusRequest struct {
	Status string `json:"status"`
}
