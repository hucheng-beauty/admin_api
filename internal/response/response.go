package response

type File struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type UpdatePassword struct {
	Id string `json:"id"`
}
