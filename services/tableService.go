package service

import (
	"GetGroundTask/models"
	"GetGroundTask/requests"
)

func AddGuests(request requests.AddGuestRequest, name string) error {
	err := models.UpdateTable(request.TableId, request.AccompanyingGuests, name)
	if err != nil {
		return err
	}
	return nil
}

func GetGuests() (*[]models.TableInfo, error) {
	tablesInfo := []models.TableInfo{}
	err := models.GetGuests(&tablesInfo)
	if err != nil {
		return nil, err
	}
	return tablesInfo, nil
}

func AddGuestsArrivals(request requests.AddGuestRequest, name string) error {
	err := models.UpdateTableArrival(request.TableId, request.AccompanyingGuests, name)
	if err != nil {
		return err
	}
	return nil
}

func GetGuestCurrent() (*[]models.GuestInfo, error) {
	guestInfo := []models.GuestInfo{}
	err := models.GetGuests(&guestInfo)
	if err != nil {
		return nil, err
	}
	return guestInfo, nil
}

func GetEmptySeats() (int64, error) {
	seats, err := models.GetEmptySeats()
	if err != nil {
		return 0, err
	}
	return seats, nil
}

func DeleteGuest(name string) error {
	err := models.DeleteGuest()
	if err != nil {
		return err
	}
	return nil
}
