package requests

type ArriveGuestRequest struct {
	AccompanyingGuests int64 `json:"accompanying_guests"`
}
