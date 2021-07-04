package requests

type EditGuestList struct {
	Table              int64 `json:"table"`
	AccompanyingGuests int64 `json:"accompanying_guests"`
}
