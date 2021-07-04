package requests

type AddGuestRequest struct {
	Table              int64 `json:"table"`
	AccompanyingGuests int64 `json:"accompanying_guests"`
}
