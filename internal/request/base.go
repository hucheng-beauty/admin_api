package request

// Pagination 页码
type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// Validate also can use gin Validator
func (p *Pagination) Validate() {
	// default offset eq zero and limit eq ten.
	if p.Offset <= 0 {
		p.Offset = 0
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
}
