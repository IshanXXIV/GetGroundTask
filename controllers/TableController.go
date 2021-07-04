package controllers

import (
	"net/http"

	"GetGroundTask/requests"
	"GetGroundTask/service"

	"github.com/labstack/echo"
)

func AddGuest(c echo.Context) error {
	req := requests.AddGuestRequest{}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	name := c.Param("name")

	errService := service.AddGuest(req, name)

	if errService != nil {
		return c.JSON(http.StatusBadRequest, errService)
	}
	return c.JSON(http.StatusOK, "Guest List Successfully Updated")
}

func GetGuests(c echo.Context) error {
	guestList, err := service.GetGuests()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, guestList)
}

func AddGuestsArrivals(c echo.Context) error {
	req := requests.ArriveGuestRequest{}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	name := c.Param("name")

	errService := service.AddGuestsArrivals(req, name)

	if errService != nil {
		return c.JSON(http.StatusBadRequest, errService)
	}
	return c.JSON(http.StatusOK, "Guest List Successfully Updated")
}

func GetGuestCurrent(c echo.Context) error {
	guestList, err := service.GetGuests()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, guestList)
}

func GetEmptySeats(c echo.Context) error {
	seats, err := service.GetEmptySeats()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, seats)
}

func DeleteGuest(c echo.Context) error {
	name := c.Param("name")
	err := service.DeleteGuest(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "Guest successfully deleted")
}
