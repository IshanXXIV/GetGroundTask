package models

type Guests struct {
	Id           int64  `json:"id"`
	TableId      int64  `json:"table_id"`
	Accompanying int64  `json:"accompanying"`
	ArrivalTime  string `json:"arrival_time"`
}
