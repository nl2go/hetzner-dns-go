package models

type Meta struct {
	Pagination struct {
		LastPage     int `json:"last_page"`
		NextPage     int `json:"next_page"`
		Page         int `json:"page"`
		PerPage      int `json:"per_page"`
		PreviousPage int `json:"previous_page"`
		TotalEntries int `json:"total_entries"`
	} `json:"pagination"`
}
