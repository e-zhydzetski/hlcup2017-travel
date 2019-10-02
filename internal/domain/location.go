package domain

type Location struct {
	ID       uint32 `json:"id"` // unique
	Place    string `json:"place"`
	Country  string `json:"country"` // len <= 50
	City     string `json:"city"`    // len <= 50
	Distance uint32 `json:"distance"`
}
