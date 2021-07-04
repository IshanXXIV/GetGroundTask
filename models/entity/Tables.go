package models

type Tables struct {
	Id        int64  `json:"id"`
	Capacity  int64  `json:"capacity"`
	Fill      int64  `json:"fill"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
