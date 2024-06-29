package dtos

type TodoRequest struct {
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
	Priority    int    `json:"priority"`
}
