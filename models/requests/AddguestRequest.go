package requests

type AddGuestRequest struct {
	TableId            int64 `json:"table"`
	AccompanyingGuests int64 `json:"accompanying_guests"`
}
