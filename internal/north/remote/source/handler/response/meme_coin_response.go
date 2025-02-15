package response

type CreateMemeCoin struct {
	ID string `json:"id"`
}

type GetMemeCoin struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	PopularityScore int    `json:"popularity_score"`
	CreatedAt       string `json:"created_at"`
}
